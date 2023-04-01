# LL! SIF Private Server

LoveLive! 学园偶像祭自用私服

## How to use?

1. 首先你需要一台电脑，在电脑上安装安卓模拟器（例如雷电模拟器），在模拟器中安装 SIF 国服官方客户端并登录自己的账号；
2. 进入游戏右上角 `各种设置 - 批量下载`，将所有数据下载到本地，大约 11G 左右；
3. 模拟器开启 root 权限，然后备份游戏数据 `/data/data/klb.android.lovelivecn/files` ~~以及 `/data/data/klb.android.lovelivecn/shared_prefs` 目录下的账号数据 `GameEngineActivity.xml` 和 `klb.android.lovelivecn_preferences.xml`~~；
4. 从登录开始的过程中使用 WireShark 等工具对游戏响应数据进行抓取，并尽可能地将所有功能都使用一遍，以便于给后续开发做参考；
5. 反编译官方游戏安装包，替换客户端使用的公钥为你自己的，并替换盛趣相关的接口地址为本地或者远程服务器的地址，然后重新打包签名，这样我们才能在服务端解密数据包，可参考我的 [llsif-cn-client](https://github.com/YumeMichi/llsif-cn-client)；
6. 都准备好后，现在需要一台有 root 权限的安卓手机（苹果 iOS 不了解）用于后续操作；
7. 在手机上安装修改后的客户端，然后将第 3 步中备份的数据解压到对应的目录下，修复好权限（chown）；
8. 修改游戏默认的请求地址，使用 [HonokaMiku](https://github.com/YumeMichi/HonokaMiku) （我没有 VS 用来编译 Windows 二进制文件，Windows 用户可以使用 [libhonoka](https://github.com/DarkEnergyProcessor/libhonoka)）解密游戏目录下的服务器配置文件 `/data/data/klb.android.lovelivecn/files/external/config/server_info.json`，将其中的 `prod.game1.ll.sdo.com` 都替换为自己的服务器地址，然后重新加密后替换掉相应文件并修复权限（chown）；
9. 上述步骤都操作完成后运行 `honoka-chan`（第一次运行会生成配置文件，根据自身情况进行配置）；
10. 手机运行游戏，使用账号密码登录，输入任意账号密码均可，上述步骤如果都没有问题，从登录开始的请求都会转发到 `honoka-chan`；
11. 后续即可根据第 4 步中抓取的数据包进行相应的服务端功能开发。

## 注意事项

1. ~~请备份账号数据后不要再通过官方客户端登录游戏，否则备份的账号信息会失效，后续登录需要填写验证码，该验证码通道可能会随着关服而关闭。~~
2. ~~后续操作都在手机上进行的原因是我用的雷电模拟器，从 3669 开始往后地卡片，只要放入响应包中，游戏客户端都会崩溃。个人太懒了不想试其他模拟器了所以都在手机上进行操作。~~
3. 模拟器系统尽量使用 32 位安卓 7，经测试兼容性最好，不会有 64 位的闲置一会儿闪退和低版本安卓 5.1 部分卡片加载会闪退（即原先上面的第 2 点）的问题，手机的话应该都没问题。

上述崩溃日志
<details>
<summary>展开查看</summary>
<pre>
I/DEBUG   ( 1368): *** *** *** *** *** *** *** *** *** *** *** *** *** *** *** ***
I/DEBUG   ( 1368): Build fingerprint: 'asus/android_x86/x86:5.1.1/LMY47I/8.3.19:user/release-keys'
I/DEBUG   ( 1368): Revision: '0'
I/DEBUG   ( 1368): ABI: 'x86'
I/DEBUG   ( 1368): pid: 3816, tid: 3834, name: GLThread 147  >>> klb.android.lovelivecn <<<
I/DEBUG   ( 1368): signal 4 (SIGILL), code 2 (ILL_ILLOPN), fault addr 0xa415f674
I/DEBUG   ( 1368):     eax 4f1c02df  ebx a44586d0  ecx 4f1c02df  edx a3db45e2
I/DEBUG   ( 1368):     esi a081a480  edi a0842100
I/DEBUG   ( 1368):     xcs 00000073  xds 0000007b  xes 0000007b  xfs 0000005f  xss 0000007b
I/DEBUG   ( 1368):     eip a415f674  ebp a3db4b48  esp a3db4b10  flags 00010282
I/DEBUG   ( 1368):
I/DEBUG   ( 1368): backtrace:
I/DEBUG   ( 1368):     #00 pc 002ab674  /data/app/klb.android.lovelivecn-1/lib/x86/libGame.so
I/DEBUG   ( 1368):     #01 pc 0028fef0  /data/app/klb.android.lovelivecn-1/lib/x86/libGame.so
I/DEBUG   ( 1368):     #02 pc 0032e11a  /data/app/klb.android.lovelivecn-1/lib/x86/libGame.so
I/DEBUG   ( 1368):     #03 pc 0032d8de  /data/app/klb.android.lovelivecn-1/lib/x86/libGame.so
I/DEBUG   ( 1368):     #04 pc 0029c9da  /data/app/klb.android.lovelivecn-1/lib/x86/libGame.so
I/DEBUG   ( 1368):     #05 pc 002a1a98  /data/app/klb.android.lovelivecn-1/lib/x86/libGame.so
I/DEBUG   ( 1368):     #06 pc 00391ac8  /data/app/klb.android.lovelivecn-1/lib/x86/libGame.so (app_klb_android_GameEngine_PFInterface_frameFlip+40)
I/DEBUG   ( 1368):     #07 pc 00001660  /data/app/klb.android.lovelivecn-1/lib/x86/libjniproxy.so (Java_klb_android_GameEngine_PFInterface_frameFlip+48)
I/DEBUG   ( 1368):     #08 pc 000d2a81  /system/lib/libart.so (art_quick_generic_jni_trampoline+49)
I/DEBUG   ( 1368):     #09 pc 000d0428  /system/lib/libart.so (art_quick_invoke_stub+72)
I/DEBUG   ( 1368):     #10 pc 003414a9  /system/lib/libart.so (art::mirror::ArtMethod::Invoke(art::Thread*, unsigned int*, unsigned int, art::JValue*, char const*)+201)
I/DEBUG   ( 1368):     #11 pc 0046d981  /system/lib/libart.so (artInterpreterToCompiledCodeBridge+129)
I/DEBUG   ( 1368):     #12 pc 0027669b  /system/lib/libart.so (bool art::interpreter::DoCall<false, false>(art::mirror::ArtMethod*, art::Thread*, art::ShadowFrame&, art::Instruction const*, unsigned short, art::JValue*)+427)
I/DEBUG   ( 1368):     #13 pc 004b5eff  /system/lib/libart.so (bool art::interpreter::DoInvoke<(art::InvokeType)2, false, false>(art::Thread*, art::ShadowFrame&, art::Instruction const*, unsigned short, art::JValue*)+399)
I/DEBUG   ( 1368):     #14 pc 000c23df  /system/lib/libart.so (art::JValue art::interpreter::ExecuteGotoImpl<false, false>(art::Thread*, art::MethodHelper&, art::DexFile::CodeItem const*, art::ShadowFrame&, art::JValue)+36319)
I/DEBUG   ( 1368):     #15 pc 00251756  /system/lib/libart.so (art::interpreter::EnterInterpreterFromStub(art::Thread*, art::MethodHelper&, art::DexFile::CodeItem const*, art::ShadowFrame&)+166)
I/DEBUG   ( 1368):     #16 pc 004a52d1  /system/lib/libart.so (artQuickToInterpreterBridge+737)
I/DEBUG   ( 1368):     #17 pc 000d2b12  /system/lib/libart.so (art_quick_to_interpreter_bridge+34)
I/DEBUG   ( 1368):     #18 pc 00e9251b  /data/dalvik-cache/x86/system@framework@boot.oat
</pre>
</details>

## 特别感谢

虹原翼 (yazawa@niconi.co.ni) @ [LLSIF 查卡器
](https://card.niconi.co.ni/) 的技术支持。

## 一些截图
<details>
<summary>展开查看</summary>
<img src="assets/ss01.png" />
<img src="assets/ss02.png" />
<img src="assets/ss03.png" />
</details>
