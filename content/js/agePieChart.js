/**
 * Created by gloomysw on 2016/06/15.
 */
var ageData=[
    {value:0, name:'0-15岁  '},
    {value:0, name:'16-25岁'},
    {value:0, name:'26-40岁'},
    {value:0, name:'40岁以上'}
];
var agePieOption = {
    backgroundColor:'#082740',
    legend: {
        orient: 'horizontal',
        left: 'left',
        bottom: 0,
        data: ['0-15岁  ','16-25岁','26-40岁','40岁以上'],
        textStyle:{
            color:'white',
            fontSize:11
        },
        itemHeight:8,
        left:'15px'
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
            data:ageData
        }
    ],
    color:['#4bcaff','#2a80b1', '#20bbb5', '#77fffa']
};

var agePieChart = echarts.init(document.getElementsByClassName('axon_showPieArea')[0]);
agePieChart.setOption(agePieOption);

/**
 * 年龄饼图处理方法
 * 创建人:邵炜
 * 创建时间:2016年6月17日16:37:14
 * @param heatMapHistList 数据对象数组
 */
function agePieProcess(heatMapHistList) {
    ageData.forEach(function (value,index) {
        ageData[index].value=0;
    });
    heatMapHistList.forEach(function(value){
        if (value.Age >= 0 && value.Age <= 15) {
            ageData[0].value++;
        }else if (value.Age >= 16 && value.Age <= 25) {
            ageData[1].value++;
        }else if (value.Age >= 26 && value.Age <= 40) {
            ageData[2].value++;
        }else{
            ageData[3].value++;
        }
    });
    agePieChart.setOption({
        series: [{
            data: ageData
        }]
    });
}