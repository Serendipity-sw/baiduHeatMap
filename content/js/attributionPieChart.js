/**
 * Created by gloomysw on 2016/06/15.
 */
var attrbutionPieData=[
    {value:0, name:'本地'},
    {value:0, name:'外地'},
    {value:0, name:'外国'}
];
var attributionPieOption = {
    backgroundColor:'#082740',
    legend: {
        orient: 'horizontal',
        left: 'left',
        bottom: 0,
        data: ['本地','外地','外国'],
        textStyle:{
            color:'white',
            fontSize:11
        },
        itemHeight:8,
        left:'6px'
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
                        formatter:  function(params){
                            return Math.round(params.percent)+"%";
                        }
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
            data:attrbutionPieData
        }
    ],
    color:['#4bcaff','#2a80b1', '#20bbb5', '#77fffa']
};

var attributionPieChart = echarts.init(document.getElementsByClassName('axon_showPieArea')[2]);
attributionPieChart.setOption(attributionPieOption);

/**
 * 归属地分布处理方法
 * 创建人:邵炜
 * 创建时间:2016年6月17日17:03:39
 * @param heatMapHistList 数据对象集合
 */
function attributionPieProcess(heatMapHistList) {
 attrbutionPieData.forEach(function(value,index){
     attrbutionPieData[index].value=0;
 });
    heatMapHistList.forEach(function(value){
        switch (value.Locale){
            case "1"://本地号码
                attrbutionPieData[0].value++;
                break;
            case "2"://外地号码
                attrbutionPieData[1].value++;
                break;
            default://国外
                attrbutionPieData[2].value++;
                break;
        }
    });
    attributionPieChart.setOption({
        series: [{
            data: attrbutionPieData
        }]
    });
}