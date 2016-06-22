/**
 * Created by gloomysw on 2016/06/15.
 */
var arpuData=[
    {value:0, name:'0-49元      '},
    {value:0, name:'50-99元'},
    {value:0, name:'100-199元'},
    {value:0, name:'200元以上'}
];
var arpuPieOption = {
    backgroundColor:'#082740',
    legend: {
        orient: 'horizontal',
        left: 'left',
        bottom: 0,
        data: ['0-49元      ','50-99元','100-199元','200元以上'],
        textStyle:{
            color:'white',
            fontSize:11
        },
        itemHeight:8,
        left:'12px'
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
                        formatter:'{d}%'  
                        //     function(params){
                        //     return Math.round(params.percent)+"%";
                        // }
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
            data:arpuData
        }
    ],
    color:['#4bcaff','#2a80b1', '#20bbb5', '#77fffa']
};

var arpuPieChart = echarts.init(document.getElementsByClassName('axon_showPieArea')[3]);
arpuPieChart.setOption(arpuPieOption);

/**
 * arpu分布数据处理
 * 创建人:邵炜
 * 创建时间:2016年6月17日16:57:18
 * @param heatMapHistList 数据对象数组
 */
function arpuPieProcess(heatMapHistList) {
    arpuData.forEach(function(value,index){
        arpuData[index].value=0;
    });
    heatMapHistList.forEach(function(value){
        if (value.Arpu >= 0 && value.Arpu <= 49) {
            arpuData[0].value++
        }else if (value.Arpu>=50 && value.Arpu<=99){
            arpuData[1].value++
        }else if(value.Arpu>=100&&value.Arpu<=199){
            arpuData[2].value++
        }else{
            arpuData[3].value++
        }
    });
    arpuPieChart.setOption({
        series: [{
            data: arpuData
        }]
    });
}