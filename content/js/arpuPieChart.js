/**
 * Created by gloomysw on 2016/06/15.
 */
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
                {value:335, name:'0-49元      '},
                {value:310, name:'50-99元'},
                {value:234, name:'100-199元'},
                {value:234, name:'200元以上'}
            ]
        }
    ],
    color:['#4bcaff','#2a80b1', '#20bbb5', '#77fffa']
};

var arpuPieChart = echarts.init(document.getElementsByClassName('axon_showPieArea')[3]);
arpuPieChart.setOption(arpuPieOption);