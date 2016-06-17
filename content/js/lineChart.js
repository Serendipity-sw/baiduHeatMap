/**
 * Created by wei on 2016/06/15.
 */
var data = [];

var lineOption = {
    tooltip: {
        trigger: 'axis',
        formatter: function (params) {
            params = params[0];
            return params.name + ' | ' + params.value[1];
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
    width:'371px'
};
var myChart = echarts.init(document.getElementsByClassName('axon_graph')[0]);
myChart.setOption(lineOption,true);

/**
 * 折线图处理方法
 * 创建人:邵炜
 * 创建时间:2016年6月17日16:25:31
 * @param heatMapHistList 数据对象数组
 */
function lineChartProcess(heatMapHistList) {
var now=new Date();
    var timeStr=now.getHours()+":"+now.getMinutes();
    data.push({"name":timeStr,"value":[timeStr,heatMapHistList.length]});
    if (data.length > 5) {
        data.shift();
    }
    myChart.setOption({
        series: [{
            data: data
        }]
    });
}