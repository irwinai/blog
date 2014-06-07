(function(){

  var kernelCache = {};

  function sum(a) {
    var s = 0;
    for (var i = 1 - a.length; i < a.length; i++) {
      s = s + a[Math.abs(i)];
    }
    return s;
  }

  function buildKernel(size) {
    // Blur of size 1 or less will result in the very same pixel.
    // So there's no point calculating kernel for it.
    if (size < 1) {
      size = 1;
      if (!kernelCache[size]) {
        kernelCache[size] = {
          kernel: [1],
          sum: 1
        }
      }
    }

    if (!kernelCache[size]) {
      var sigma = size / 3;
      var ss = sigma * sigma;
      var factor = 2 * Math.PI * ss;
      kernel = [];
      for (var i = 0; i < size; i++) {
        kernel[i] = Math.exp(-(i * i) / (2 * ss)) / Math.sqrt(factor);
      }
      kernelCache[size] = {
        kernel: kernel,
        sum: sum(kernel)
      }
    }

    return kernelCache[size];
  }

  function createImageData(imgData) {
    return document.createElement('canvas').getContext('2d').createImageData(imgData.width, imgData.height);
  }

  function verticalBlur(imgData, radiusFunc) {
    var width = imgData.width, height = imgData.height;

    var newImgData = createImageData(imgData);

    var r = false, new_r;

    for (var x = 0; x < width; x++) {
      for (var y = 0; y < height; y++) {
        var new_r = radiusFunc(x, y);
        if (new_r < 1)
          new_r = 1;
        if (r !== new_r) {
          r = new_r;
          var k = buildKernel(r);
          var kernel = k.kernel;
          var kernelSize = kernel.length;
          var kernelSum = k.sum;
        }

        var r = 0, g = 0, b = 0, a = 0;
        var j, i, n2;
        var n = 4 * (y * width + x);

        if (kernelSize <= 1) {
          // Just copy pixel when no blur needed
          for (j = 0; j < 4; j++) {
            newImgData.data[n + j] = imgData.data[n + j];
          }
        }
        else {
          for (j = 1 - kernelSize; j < kernelSize ; j++) {
            i = Math.abs(j);

            if (y + j < 0 || y + j >= height) {
              // for edges use transparent self
              r += imgData.data[n + 0] * kernel[i];
              g += imgData.data[n + 1] * kernel[i];
              b += imgData.data[n + 2] * kernel[i];
              a += imgData.data[n + 3] * kernel[i];
            }
            else {
              n2 = 4 * ((y + j) * width + x);
              r += imgData.data[n2 + 0] * kernel[i];
              g += imgData.data[n2 + 1] * kernel[i];
              b += imgData.data[n2 + 2] * kernel[i];
              a += imgData.data[n2 + 3] * kernel[i];
            }
          }
          newImgData.data[n + 0] = r / kernelSum;
          newImgData.data[n + 1] = g / kernelSum;
          newImgData.data[n + 2] = b / kernelSum;
          newImgData.data[n + 3] = a / kernelSum;
        }
      }
    }

    return newImgData;
  }


  function horizontalBlur(imgData, radiusFunc) {
    var width = imgData.width, height = imgData.height;

    var newImgData = createImageData(imgData);

    var r = false, new_r;

    for (var x = 0; x < width; x++) {
      for (var y = 0; y < height; y++) {
        var new_r = radiusFunc(x, y);
        if (new_r < 1)
          new_r = 1;
        if (r !== new_r) {
          r = new_r;
          var k = buildKernel(r);
          var kernel = k.kernel;
          var kernelSize = kernel.length;
          var kernelSum = k.sum;
        }

        var r = 0, g = 0, b = 0, a = 0;
        var j, i, n2;
        var n = 4 * (y * width + x);

        if (kernelSize <= 1) {
          for (j = 0; j < 4; j++) {
            newImgData.data[n + j] = imgData.data[n + j];
          }
        }
        else {
          for (j = 1 - kernelSize; j < kernelSize ; j++) {
            i = Math.abs(j);

            if (x + j < 0 || x + j >= width) {
              r += imgData.data[n + 0] * kernel[i];
              g += imgData.data[n + 1] * kernel[i];
              b += imgData.data[n + 2] * kernel[i];
              a += imgData.data[n + 3] * kernel[i];
            }
            else {
              n2 = 4 * (y * width + x + j);
              r += imgData.data[n2 + 0] * kernel[i];
              g += imgData.data[n2 + 1] * kernel[i];
              b += imgData.data[n2 + 2] * kernel[i];
              a += imgData.data[n2 + 3] * kernel[i];
            }
          }
          newImgData.data[n + 0] = r / kernelSum;
          newImgData.data[n + 1] = g / kernelSum;
          newImgData.data[n + 2] = b / kernelSum;
          newImgData.data[n + 3] = a / kernelSum;
        }
      }
    }

    return newImgData;
  }

  window.varBlur = function(canvas, radiusFunc) {
    var imgData = canvas.getImageData(0, 0, canvas.canvas.width, canvas.canvas.height);
    console.log(imgData)
    var newImgData = verticalBlur(imgData, radiusFunc);
    canvas.putImageData(horizontalBlur(newImgData, radiusFunc), 0, 0);
  }

})()