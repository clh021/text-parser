# text-parser
text parser by configured


    字符串输入后，根据规则，被处理后存储到处理中心的对象内部中，
    可能被截取为数组，也可能存为一个字符串
    处理完成后，由最外层程序取出最后的 lastStr 输出

    如果处理结果中一直是 lastArr(lastStr为空),则自动将其 JSON 为 lastStr

本项目最初是为硬件信息识别开始，硬件信息的命令方式从config.json中已经看出很多已经支持了JSON输出

而 dmidecode 方式可以参考 [项目一](https://github.com/xemoe/dmid2json)，[项目二](https://github.com/paulstuart/dmijson) 

所以本项目可能暂时不会继续

```shell
# 参数解析规则
# 解析 hostnamectl 的输出
text-parser "file:filename.text,split:\n,trim,split::,trim,json:encode"
# 解析 html 的元素值
text-parser "url:http://t.weather.sojson.com/api/weather/city/101200101,split:\n"
text-parser "file:1.html,split:\n"
# 解析 lscpu 的输出
text-parser "file:filename.text,split:\n"
# 解析 lshw 的输出
text-parser
```

# conf

## methods

```javascript
return {
    "EQ": "等于", // EQUAL 等于
    "NE": "不等于", // NOT EQUAL 不等于
    "GT": "大于", // GREATER THAN 大于
    "LT": "小于", // LESS THAN 小于
    "GE": "大于等于", // GREATER THAN OR EQUAL 大于等于
    "LE": "小于等于", // LESS THAN OR EQUAL 小于等于
    "SC": "包含字符", // String Contain 包含字符
    //TODO: 从 hardinfo 获取结果
}
```

## autoList

- 每一个函数解决一个问题
- 每一次处理不能破坏数据本身

```javascript
window.parseHardinfoTxt=(hardinfoTxt) => {
  let resultObj = {};
  // getTitleFunc : (areaStr) => {title:titleStr, content: contentStr};
  let fixTitleToMap = (arr, getTitleFunc) => {
    let map = {};
    for (let index = 1; index < arr.length; index++) {
      let prevArr = arr[index-1].split("\n");
      let mapEle = getTitleFunc(prevArr[prevArr.length - 1] +"\n"+ arr[index]);
      map[mapEle.title] = mapEle.content;
    }
    return map;
  }
  let txtToMap = (txt, splitStr, getTitleFunc) => {
    return fixTitleToMap(txt.split(splitStr), getTitleFunc);
  }

  // 末尾增加多个换行以帮助解析去除块与块的标题间隙
  let hardinfoTxtWithLastfix = hardinfoTxt+"\n\n\n\n\n\n";

  // 根据 下一行是 ***** 来断定一级标题
  resultObj = txtToMap(hardinfoTxtWithLastfix, "\n*", (areaStr) => {
    let lineArr = areaStr.split("\n")
      .filter(e => e.length > 0)
      .filter(e => !e.startsWith("*"));
    lineArr.pop(1);// 移除最后一个元素 下一个块标题
    return {title: lineArr[0], content:lineArr.slice(1).join("\n")}
  });

  // 根据 下一行是 ----- 来断定二级标题
  for (const [title, content] of Object.entries(resultObj)) {
    resultObj[title] = txtToMap(content, "\n--", (areaStr) => {
      let lineArr = areaStr.split("\n")
        .filter(e => !e.startsWith("--"));
      lineArr.pop(1);// 移除最后一个元素 下一个块标题
      return {title: lineArr[0], content:lineArr.slice(1).join("\n")}
    })
  }

  // 根据 当前行是 -[title]- 来断定三级标题
  for (const [title, titleObj] of Object.entries(resultObj)) {
    for (const [subtitle, subtitleObj] of Object.entries(titleObj)) {
      // debugger
      resultObj[title][subtitle] = txtToMap(subtitleObj, "-\n", (areaStr) => {
        let lineArr = areaStr.split("\n");
        lineArr.pop(1);// 移除最后一个元素 下一个块标题
        return {title: lineArr[0].substring(1), content:lineArr.slice(1).join("\n")}
      })
    }
  }

  // 进一步数据分析应该在具体使用数据的模块来处理:
  // - 各个模块之间的数据格式和层级都不一样
  // - 部分数据的 key:val 中 key 作为表格元素有重复，JSON 处理会丢失数据

  return resultObj;
}

function getDetectionValueByPipeMethod(method, data, args) {
    let shouldType = "";
    let val = data;
    switch (method) {
        case "split":
            if (typeof val === 'string') {
                val = val.split(args[0]);
            } else {
                shouldType="string";
            }
            break;
        case "contain":
            if (Array.isArray(val)) {
                val = val.filter(e => e.indexOf(args[0]) != -1)
            } else {
                shouldType="array";
            }
            console.log("contain.val:", val);
            break;
        case "join":
            if (Array.isArray(val)) {
                val = val.join(args[0]);
            } else {
                shouldType="array";
            }
            break;
        default:
            break;
    }
    if (shouldType) {
        console.warn(`DetectPipe:${method}: type should "${shouldType}" but "${typeof val}"`, val);
    }
    return val;
}

/**
 * 根据自动核验数据来源 核验数据
 * @param {String} content
 * @param {Array} pipes config
 * @returns {String} value
 */
function getDetectionValueByPipe(content, pipes) {
    let val = content;
    for (const e of pipes) {
        let method = e[0];
        let args = e.slice(1);
        val = getDetectionValueByPipeMethod(method, val, args);
    }
    return val;
}
```