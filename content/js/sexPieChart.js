/**
 * Created by gloomysw on 2016/06/15.
 */
var sexData=[
    {value:0, name:'男性'},
    {value:0, name:'女性'}
];
var sexPieOption = {
    backgroundColor:'#082740',
    legend: {
        orient: 'horizontal',
        left: 'left',
        bottom: 0,
        data: ['男性','女性'],
        textStyle:{
            color:'white',
            fontSize:11
        },
        itemHeight:8,
        left:'44px'
    },
    series : [
        {
            name: '访问来源',
            type: 'pie',
            radius : '60%',
            center: ['50%', '39%'],
            itemStyle:{
                normal:{
                    label:{
                        show: true,
                        formatter: '{d}%'
                    },
                    labelLine :{show:true}
                }
            },
            label: {
                normal: {
                    show: true,
                    textStyle:{
                        color:'white',
                        fontSize:10
                    }
                },
                emphasis: {
                    show: true,
                    textStyle: {
                        fontSize: '10',
                        fontWeight: 'bold'
                    }
                }
            },
            labelLine:{
                normal:{
                    show:false,
                    length:2,
                    length2:2
                }
            },
            data:sexData
        }
    ],
    color:['#4bcaff','#2a80b1', '#20bbb5', '#77fffa']
};

var sexPieChart = echarts.init(document.getElementsByClassName('axon_showPieArea')[1]);
sexPieChart.setOption(sexPieOption);

/**
 * 性别分布处理方法
 * 创建人:邵炜
 * 创建时间:2016年6月17日17:18:17
 * @param heatMapHistList 数据对象集合
 */
function sexPieProcess(heatMapHistList) {
    sexData.forEach(function(index){
        sexData[index].value=0;
    });
    heatMapHistList.forEach(function(index,value){
        switch (value.Gender){
            case "0"://女
                sexData[1].value++;
                break;
            default:
                sexData[0].value++;
                break;
        }
    });
}