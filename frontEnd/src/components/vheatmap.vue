<template>
  <v-container fluid>
    <v-layout row v-if="!tableView">
      <v-flex class="calender-map">
        <svg viewBox="0 0 700 135" xmlns="http://www.w3.org/2000/svg" v-if="CalendarData.year">
          <text transform="translate(350,10)" font-size="10px"  font="sans-serif" width="100%" style="text-anchor: middle;" fill="black">{{CalendarData.title}}</text>
          <g transform="translate(40,30)" 
          font-size="10px" font="sans-serif" shape-rendering="crispEdges" width="100%">
          <text transform="translate(-25,42)rotate(-90)" style="text-anchor: middle;  fill: white;">{{CalendarData.year}}</text>

          <text class="dow" v-for="d in 7" :transform="getDayTranslations(d)"  dy="-.25em"     fill="black" text-anchor="end">{{dayNames[d-1]}}</text>

          <g v-for="i in 12" class="month" :transform="getTranslation(i)" >
            <text style="text-anchor: end;" dy="-.25em">{{monthNames[i-1]}}</text>
          </g>
          <rect  v-for="d in AllDays" 
          stroke="#666"
          text-anchor="end"
          :width="cSize" :height="cSize" 
          :x="getX(d)" :y="getY(d)" :fill="getColor(d)">
        </rect>

        <path v-for="m in 12" :d="getPath(m)"
        class="border" fill="none" stroke="#000" stroke-width="2px"></path>
        <g v-for="(pl, ind) in CalendarData.pcounts">
        <rect stroke="#666" :width="cSize" :height="cSize" 
          :x="getLegendx(ind)+'em'" y="9em" :fill="colors[pl.cat_id-1]"> </rect>
          <text :x="getLegendx(ind)+2+'em'" dy="10em" font-size="10px" fill="black">{{pl.cat}}</text>
        </g>
        <g v-for="(ul, ind) in CalendarData.ucounts">
        <rect  stroke="#666"  :width="cSize" :height="cSize" 
          :x="getLegendx(ind,false)+'em'" y="9em" :fill="colors[ul.cat_id-1]"> </rect> 
          <text :x="getLegendx(ind,false)+ 2 +'em'" dy="10em" fill="black">{{ul.cat}}</text>
        </g>
      </g>
    </svg>
  </v-flex>
  <v-flex>
    
  </v-flex>
</v-layout>
<v-layout align-top justify-center v-if="tableView" class="tr pop-up-message">
  <v-data-table
  :headers="headers"
  :items="CalendarData.pcounts"
  hide-actions
  class="elevation-1"
  >
  <template slot="items" slot-scope="props">
    <td class="text-xs-right">{{ props.item.name }}</td>
    <td class="text-xs-right">{{ props.item.cat }}</td>
    <td class="text-xs-right">{{ props.item.count }}</td>
    <td class="text-xs-right">{{ props.item.count / CalendarData.total | percent }}</td>
  </template>
</v-data-table>
<v-data-table
:headers="headers"
:items="CalendarData.ucounts"
hide-actions
class="elevation-1"
>
<template slot="items" slot-scope="props">
  <td class="text-xs-right">{{ props.item.name }}</td>
  <td class="text-xs-right">{{ props.item.cat }}</td>
  <td class="text-xs-right">{{ props.item.count }}</td>
  <td class="text-xs-right">{{ props.item.count / CalendarData.total | percent }}</td>
</template>
</v-data-table>
</v-layout>
<v-layout align-end justify-end>
<!--   <v-flex >
    <div v-for="c in colors" style="height:1.4em;  width: 1.4em; border-color: #00E676; margin:2px; background:coral;" v-bind:style="{background:c}">
      {{c}}
    </div>
  </v-flex>  -->
  <v-btn color="red" outline v-on:click="flipView()">Stats</v-btn>
  <v-btn color="info"   v-on:click="getSvg(CalendarData.file_title)">Export</v-btn>
</v-layout>
</v-container>
</template>

<script>
import * as d3 from "d3";
import moment from "moment";
import numeral from "numeral";

export default {
  data: () => ({
    legend: {},
    test: "hello",
    AllDays: [],
    cSize :12,
    heatColor: "#002b53",
    colorRange: null,
    firstSunday: null,
    dayNames:['Sun','Mon','Tue','Wed','Thu','Fri','Sat'],
    monthNames : [ "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec" ],
    tableView : false,
    headers:[
    { text: 'Leave Type', value: 'name' },
    { text: 'Leave Category', value: 'cat' },
    { text: 'Count', value: 'count' },
    { text: 'Percentage', value: '' },
    ],
    colors : ['#BF360C','#FFFFFF','#FFD600','#00E676','#2E7D32','#9E9D24','#C6FF00','#00B0FF','#039BE5','#7C4DFF','#8C9EFF','#B39DDB','#FF4081','#AB47BC']

  }),
  methods: {
    flipView: function(){
      this.tableView = !this.tableView;
    },
    getLegendx: function(index,planned=true){
      // var sum = 0;
      // var target = this.CalendarData.pcounts;
      // var i;
      // if (!planned){
      //   for(i=0; i < target.length; i++){
      //     sum +=  (2 + target[i].cat.length);
      //   }  
      //   target = this.CalendarData.ucounts;
      // }

      // for(i=0; i < index; i++){
      //   sum +=  (2 + target[i].cat.length);
      // }
      var spacer = 14;
      if (planned)
        return index * spacer;
      else
        return (this.CalendarData.pcounts.length  + index) *spacer;
      // var space = sum/1.2;
      // console.log("====>",index,this.CalendarData.pcounts[i].cat, sum, space);
      // return space;
    },
    getSvg: function(fileName){
      var content = this.$el.firstChild.firstChild.innerHTML;
      const blob = new Blob([content], {type: 'image/svg+xml'})
      const e = document.createEvent('MouseEvents'),
      a = document.createElement('a');
      a.download = fileName + ".svg";
      a.href = window.URL.createObjectURL(blob);
      a.dataset.downloadurl = ['image/svg+xml', a.download, a.href].join(':');
      e.initEvent('click', true, false, window, 0, 0, 0, 0, 0, false, false, false, false, 0, null);
      a.dispatchEvent(e);
    },
    weekNo: function(daDate){
      return Math.ceil((((daDate - this.firstSunday) / 86400000) + 1)/7);
    },
    getX: function(daDate){
      return this.cSize * this.weekNo(daDate);
    },
    getY: function(daDate){
      return this.cSize * daDate.getDay();
    },
    getColor: function(daDate) {
      var kw = moment(daDate).format("DD-MM");
      var z = 0.0;
      if ( this.CalendarData.days == undefined){
        return "#424242";
      }

      if (!( kw in  this.CalendarData.days )){
        return "#424242";
      }
      return this.colors[this.CalendarData.days[kw].cat_id-1];
    },
    getValue: function(d){
      return "hello";
    },
    getTranslation: function(i){
      return "translate("+ (8+ i* 4.3 * this.cSize) + ",-3)";
    },
    getDayTranslations: function(d){
      return "translate(0,"+ d * this.cSize + ")";
    },
    getPath: function(m) {
      var t0 = new Date(Date.UTC(this.CalendarData.year, m - 1 , 1));
      var t1 = new Date(Date.UTC(this.CalendarData.year, m , 0));
      var d0 = t0.getDay();
      var d1 = t1.getDay();
      var w0 = this.weekNo(t0);
      var w1 = this.weekNo(t1);
      return "M" + (w0 + 1) * this.cSize + "," + d0 * this.cSize
      + "H" + w0 * this.cSize + "V" + 7 * this.cSize
      + "H" + w1 * this.cSize + "V" + (d1 + 1) * this.cSize
      + "H" + (w1 + 1) * this.cSize + "V" + 0
      + "H" + (w0 + 1) * this.cSize + "Z";
    },
    generateDates: function() {
      var fromDate = new Date(Date.UTC(this.CalendarData.year, 0, 1));
      var ToDate = new Date(Date.UTC(this.CalendarData.year,11,31));
      this.AllDays.push(fromDate);
      while (fromDate < ToDate) {
        fromDate = new Date(Date.UTC(fromDate.getUTCFullYear(), fromDate.getUTCMonth(), fromDate.getUTCDate() + 1));
        this.AllDays.push(fromDate);
      }
    }
  },
  created: function () {
    // console.log("hello");
    this.generateDates();
    this.colorRange = d3.scale.linear().range(["white", this.heatColor]).domain([0, 1]);
    var fd_dow = new Date(Date.UTC(this.CalendarData.year,0,1)).getDay();
    this.firstSunday = new Date(Date.UTC(this.CalendarData.year, 0, 1 - fd_dow));
    
  },
  props:['CalendarData'],
  filters: {
    percent: function (value) {
      return numeral(value).format("0.00%"); 
    }
  }
}
</script>

