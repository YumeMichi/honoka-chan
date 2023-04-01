var Button = function() {
};

Button.TAPPED = 'tapped';

Button.initialize = function(element, onTap) {
  var touchInfo = null;

  if (typeof onTap === 'function') {
    onTap = onTap.bind(element);
  } else {
    onTap = function() {};
  }

  function d2(x1, y1, x2, y2) {
    var dx = x2 - x1, dy = y2 - y1;
    return dx * dx + dy * dy;
  }

  function startTap(touch) {
    touchInfo = {startX: touch.clientX, startY: touch.clientY, tapped: false};
    element.classList.add(Button.TAPPED);
  }

  function endTap() {
    element.classList.remove(Button.TAPPED);
    touchInfo = null;
  }

  element.addEventListener('touchstart', function(e) {
    startTap(e.changedTouches[0]);
    e.preventDefault();
  });

  element.addEventListener('touchmove', function(e) {
    if (touchInfo) {
      var x0 = touchInfo.startX;
      var y0 = touchInfo.startY;
      var x = e.changedTouches[0].clientX;
      var y = e.changedTouches[0].clientY;

      var current = document.elementFromPoint(x, y);

      const THRESHOLD_SQ_PX = 50 * 50;
      if (!element.contains(current) || d2(x0, y0, x, y) > THRESHOLD_SQ_PX) {
        endTap();
      }
    }

    e.preventDefault();
  });

  element.addEventListener('touchend', function(e) {
    if (touchInfo && !touchInfo.tapped) {
      touchInfo.tapped = true;
      onTap();
    }

    endTap();
    e.preventDefault();
  });

  element.addEventListener('touchcancel', function(e) {
    endTap();
    e.preventDefault();
  });

  return element;
};
