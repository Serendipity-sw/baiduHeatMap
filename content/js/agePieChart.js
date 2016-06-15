/**
 * Created by gloomysw on 2016/06/15.
 */
var agePieOption = {
    legend: {
        orient: 'vertical',
        left: 'left',
        data: ['直接访问','邮件营销','联盟广告','视频广告','搜索引擎']
    },
    backgroundColor:'#082740',
    series : [
        {
            name: '访问来源',
            type: 'pie',
            radius : '45%',
            center: ['50%', '50%'],
            label: {
                normal: {
                    show: true,
                    position: 'outside'
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
                  show:false
              }
            },
            data:[
                {value:335, name:'22%'},
                {value:310, name:'22%'},
                {value:234, name:'22%'},
                {value:135, name:'22%'},
                {value:1548, name:'22%'}
            ]
        }
    ]
};

var agePieChart = echarts.init(document.getElementsByClassName('axon_showPieArea')[0]);
agePieChart.setOption(agePieOption);