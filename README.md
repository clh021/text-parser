# text-parser
text parser by configured

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

```javascript
return [{
    class:"CPU 规格",
    name:"CPU 型号",
    matchDes:"Inter酷睿处理器",
    matchMethod:"SC",
    matchVal:"Intel(R) Core(TM)",
    dataConf:{
        type:"cmd",
        cmd:`cat /proc/cpuinfo | grep "model name" | cut -f2 -d: | uniq -c`,
    },
    checked: null,
    optional: false,
    val: null
},{
    class:"CPU 规格",
    name:"CPU 核心数",
    matchDes:"至少四核心",
    matchMethod:"GE",
    matchVal:4,
    dataConf:{
        type:"javascript",
        cmd:`navigator.hardwareConcurrency`,
    },
    checked: null,
    optional: false,
    val: null
},{
    class:"CPU 规格",
    name:"CPU 线程数",
    matchDes:"至少四线程",
    matchMethod:"GE",
    matchVal:4,
    dataConf:{
        type:"cmd",
        cmd:`cat /proc/cpuinfo |grep processor|wc -l`,
    },
    checked: null,
    optional: false,
    val: null
},{
    class:"内存规格",
    name:"内存配置容量",
    matchDes:"内存至少8G",
    matchMethod:"GE",
    matchVal:8,
    dataConf:{
        type:"cmd",
        cmd:`lsmem -b | grep online | awk '{print $4/1024/1024/1024}'`,
        pipes: [],
    },
    checked: null,
    optional: false,
    val: null
},{
    class:"内存规格",
    name:"内存配置容量",
    matchDes:"内存至少8G",
    matchMethod:"GE",
    matchVal:8,
    dataConf:{
        type:"cmd",
        cmd:`lsmem -b | grep online | awk '{print $4/1024/1024/1024}'`,
        pipes: {},
    },
    checked: null,
    optional: false,
    val: null
},{
    class:"显示规格",
    name:"显示器分辨率",
    matchDes:"显示器分辨率",
    matchMethod:"SC",
    matchVal:"1920x1080",
    dataConf:{
        type:"hardinfoObj",
        cmd:`Computer.Summary.Display`,
        pipes: [
            ["split","\n"],
            ["contain","Resolution"],
            ["join","\n"],
            ["split",":"],
            ["SC","x"],
            ["join","\n"],
        ],
    },
    checked: null,
    optional: false,
    val: null
},{
    class:"内存规格",
    name:"内存配置容量",
    matchDes:"内存至少18G",
    matchMethod:"GE",
    matchVal:18,
    dataConf:{
        type:"cmd",
        cmd:`lsmem -b | grep online | awk '{print $4/1024/1024/1024}'`,
    },
    checked: null,
    optional: false,
    val: null
}];
```

```javascript
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