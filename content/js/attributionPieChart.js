/**
 * Created by gloomysw on 2016/06/15.
 */
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
                {value:335, name:'本地'},
                {value:310, name:'外地'},
                {value:234, name:'外国'}
            ]
        }
    ],
    color:['#4bcaff','#2a80b1', '#20bbb5', '#77fffa']
};

var attributionPieChart = echarts.init(document.getElementsByClassName('axon_showPieArea')[2]);
attributionPieChart.setOption(attributionPieOption);