package handler

import (
	"encoding/json"
	"honoka-chan/model"
	"honoka-chan/utils"
	"net/http"
	"path"
	"time"

	"github.com/forgoer/openssl"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type ErrMsg struct {
	Error int    `json:"error"`
	Msg   string `json:"msg"`
}

func WebLogin(ctx *gin.Context) {
	area := ctx.PostForm("area")
	user := ctx.PostForm("user")
	pass := ctx.PostForm("pass")
	if area == "" || user == "" || pass == "" {
		ctx.JSON(http.StatusOK, model.Msg{
			Code:     1,
			Message:  "参数不完整！",
			Redirect: "",
		})
		return
	}

	userName := " " + area + "-" + user
	var userId int
	exists, err := UserEng.Table("users").Where("phone = ? AND password = ?", userName, openssl.Md5ToString(pass)).Cols("userid").Get(&userId)
	CheckErr(err)
	if !exists {
		ctx.JSON(http.StatusOK, model.Msg{
			Code:     1,
			Message:  "账号不存在或者密码有误！",
			Redirect: "",
		})
		return
	}

	session := sessions.Default(ctx)
	session.Options(sessions.Options{
		MaxAge: 3600 * 24,
	})
	session.Set("userid", userId)
	session.Save()

	ctx.JSON(http.StatusOK, model.Msg{
		Code:     0,
		Message:  "登录成功！",
		Redirect: "/admin/index",
	})
}

func WebLogout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Options(sessions.Options{
		Path:   "/admin",
		MaxAge: -1,
	})
	session.Save()

	ctx.Redirect(http.StatusFound, "/admin/login")
}

func Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	CheckErr(err)

	tmpPath := path.Join("./temp", file.Filename)
	err = ctx.SaveUploadedFile(file, tmpPath)
	CheckErr(err)

	data := model.Data{}
	err = json.Unmarshal([]byte(utils.ReadAllText(tmpPath)), &data)
	CheckErr(err)

	session := UserEng.NewSession()
	defer session.Close()
	if err = session.Begin(); err != nil {
		session.Rollback()
		panic(err)
	}

	for _, team := range data.Team {
		if team.Cardid == 0 {
			continue
		}
		var unitId, unitExp, unitRarity, unitHp, unitSigned int
		exists, err := MainEng.Table("common_unit_m").Join("LEFT", "unit_m", "common_unit_m.unit_id = unit_m.unit_id").
			Where("unit_m.unit_number = ?", team.Cardid).
			Cols("common_unit_m.unit_id,common_unit_m.exp,unit_m.rarity,common_unit_m.max_hp,common_unit_m.is_signed").
			Get(&unitId, &unitExp, &unitRarity, &unitHp, &unitSigned)
		CheckErr(err)

		if !exists {
			session.Rollback()
			ctx.JSON(http.StatusOK, ErrMsg{Error: 1, Msg: "卡片不存在！"})
			return
		}

		if unitRarity != 4 {
			session.Rollback()
			ctx.JSON(http.StatusOK, ErrMsg{Error: 1, Msg: "仅支持导入UR卡片！"})
			return
		}

		var diffExp, diffSmile, diffPure, diffCool int
		_, err = MainEng.Table("unit_level_limit_pattern_m").Where("unit_level_limit_id = 1 AND unit_level = 350").
			Cols("next_exp,smile_diff,pure_diff,cool_diff").Get(&diffExp, &diffSmile, &diffPure, &diffCool)
		CheckErr(err)

		isSigned := false
		if unitSigned == 1 {
			isSigned = true
		}

		var skillExp int
		if team.Skilllevel != 8 {
			skillExp = 0
		} else {
			skillExp = 29900
		}

		unitData := model.UnitData{
			UserID:                      ctx.GetInt("userid"),
			UnitID:                      unitId,
			Exp:                         unitExp + diffExp,
			NextExp:                     0,
			Level:                       350,
			MaxLevel:                    350,
			LevelLimitID:                1,
			Rank:                        2,
			MaxRank:                     2,
			Love:                        1000,
			MaxLove:                     1000,
			UnitSkillExp:                skillExp,
			UnitSkillLevel:              team.Skilllevel,
			MaxHp:                       unitHp,
			UnitRemovableSkillCapacity:  8,
			FavoriteFlag:                false,
			DisplayRank:                 2,
			IsRankMax:                   true,
			IsLoveMax:                   true,
			IsLevelMax:                  true,
			IsSigned:                    isSigned,
			IsSkillLevelMax:             true,
			IsRemovableSkillCapacityMax: true,
			InsertDate:                  time.Now().Format("2006-01-02 03:04:05"),
		}

		_, err = session.Table("user_unit_m").Insert(&unitData)
		if err != nil {
			session.Rollback()
			panic(err)
		}

		if err = session.Commit(); err != nil {
			session.Rollback()
			panic(err)
		}
	}

	ctx.JSON(http.StatusOK, ErrMsg{Error: 0, Msg: "上传成功！"})
}
