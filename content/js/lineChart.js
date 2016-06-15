/**
 * Created by wei on 2016/06/15.
 */
function randomData() {
    now = new Date(+now + oneDay);
    value = value + Math.random() * 21 - 10;
    return {
        name: now.toString(),
        value: [
            [now.getFullYear(), now.getMonth() + 1, now.getDate()].join('-'),
            Math.round(value)
        ]
    }
}

var data = [];
var now = +new Date(1997, 9, 3);
var oneDay = 24 * 3600 * 1000;
var value = Math.random() * 1000;
for (var i = 0; i < 1000; i++) {
    data.push(randomData());
}

var lineOption = {
    tooltip: {
        trigger: 'axis',
        formatter: function (params) {
            params = params[0];
            var date = new Date(params.name);
            return date.getDate() + '/' + (date.getMonth() + 1) + '/' + date.getFullYear() + ' : ' + params.value[1];
        },
        axisPointer: {
            animation: false
        }
    },
    map:{
        mapLocation:{
            x:'center',
            y:'center'
        }
    },
    xAxis: {
        type: 'time',
        splitLine: {
            show: false
        },
        axisLine:{
            lineStyle:{
                color:'#999999'
            }
        },
        axisTick:{
            inside:true
        },
        axisLabel:{
            textStyle:{
                color:'#a4aab0'
            }
        }
    },
    yAxis: {
        type: 'value',
        boundaryGap: [0, '100%'],
        splitLine: {
            show: true,
            lineStyle:{
                color:['#3b5265'],
                type:'dotted'
            }
        },
        axisLine:{
            lineStyle:{
                color:'#999999'
            }
        },
        axisTick:{
            inside:true
        },
        axisLabel:{
            textStyle:{
                color:'#a4aab0'
            }
        }
    },
    series: [{
        type: 'line',
        showSymbol: false,
        hoverAnimation: false,
        data: data,
        lineStyle:{
            normal:{
                color:'#65d5d1'
            }
        }
    }],
    backgroundColor:'#082740',
    width:'371px',
    top:'-15555px'
};
var myChart = echarts.init(document.getElementsByClassName('axon_graph')[0]);
myChart.setOption(lineOption,true);
// app.timeTicket =
setInterval(function () {

    for (var i = 0; i < 5; i++) {
        data.shift();
        data.push(randomData());
    }

    myChart.setOption({
        series: [{
            data: data
        }]
    });
}, 1000);