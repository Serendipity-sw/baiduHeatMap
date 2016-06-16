/**
 * Created by gloomysw on 2016/06/15.
 */
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
            data:[
                {value:335, name:'男性'},
                {value:310, name:'女性'}
            ]
        }
    ],
    color:['#4bcaff','#2a80b1', '#20bbb5', '#77fffa']
};

var sexPieChart = echarts.init(document.getElementsByClassName('axon_showPieArea')[1]);
sexPieChart.setOption(sexPieOption);