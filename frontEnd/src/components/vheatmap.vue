<template>
  <v-container fluid>
    Hello heatmap {{year}}
    <div class="calender-map">
      <!-- <div> -->
        <svg width="100%" data-height="0.5678" class="RdYlGn" viewBox="0 0 900 105">
          <g transform="translate(132,20)">
            <text transform="translate(-38,42)rotate(-90)" style="text-anchor: middle;  fill: white;">{{year}}</text>

            <text class="dow" v-for="d in 7" :transform="getDayTranslations(d)"  dy="-.25em">{{dayNames[d-1]}}</text>

            <g v-for="i in 12" class="month" :transform="getTranslation(i)" >
              <text style="text-anchor: end;" dy="-.25em">{{monthNames[i-1]}}</text>
            </g>
            <rect  v-for="d in AllDays" 
            class="day" 
            :width="cellSize" :height="cellSize" 
            :x="getX(d)" :y="getY(d)" :fill="getColor()">
            </rect>

            <path v-for="m in 12" class="border" id="Aug" :d="getPath(m)"></path>
          </g>
        </svg>
      </div>
    </v-container>
  </template>

  <script>
  import * as d3 from "d3";
  export default {
    data: () => ({
      test: "hello",
      AllDays: [],
      cSize :12,
      heatColor: "#002b53",
      colorRange: null,
      firstSunday: null,
      dayNames:['Sun','Mon','Tue','Wed','Thu','Fri','Sat'],
      monthNames : [ "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec" ]
    }),
    methods: {
      weekNo: function(daDate){
        return Math.ceil((((daDate - this.firstSunday) / 86400000) + 1)/7);
      },
      getX: function(daDate){
        return this.cSize * this.weekNo(daDate);
      },
      getY: function(daDate){
        return this.cSize * daDate.getDay();
      },
      getColor: function() {
        var x = Math.random();   
        return this.colorRange(x);
      },
      getValue: function(d){
        return "hello";
      },
      getTranslation: function(i){
        return "translate("+ (8+ i*52) + ",-3)";
      },
      getDayTranslations: function(d){
        return "translate(-5,"+ d * this.cSize + ")";
      },
      getPath: function(m) {
        var t0 = new Date(Date.UTC(this.year, m - 1 , 1));
        var t1 = new Date(Date.UTC(this.year, m , 0));
        var d0 = t0.getDay();
        var d1 = t1.getDay();
        var w0 = this.weekNo(t0);
        var w1 = this.weekNo(t1);
      // d0 = +day(t0), w0 = +week(t0),
      // d1 = +day(t1), w1 = +week(t1);
      return "M" + (w0 + 1) * this.cSize + "," + d0 * this.cSize
      + "H" + w0 * this.cSize + "V" + 7 * this.cSize
      + "H" + w1 * this.cSize + "V" + (d1 + 1) * this.cSize
      + "H" + (w1 + 1) * this.cSize + "V" + 0
      + "H" + (w0 + 1) * this.cSize + "Z";
    },
    generateDates: function() {
      var fromDate = new Date(Date.UTC(this.year, 0, 1));
      var ToDate = new Date(Date.UTC(this.year,11,31));
      this.AllDays.push(fromDate);
      while (fromDate < ToDate) {
        fromDate = new Date(Date.UTC(fromDate.getUTCFullYear(), fromDate.getUTCMonth(), fromDate.getUTCDate() + 1));
        this.AllDays.push(fromDate);
      }
    }
  },
  created: function () {
    if (Number.isInteger(this.cellSize)){
      this.cSize = this.cellSize;
    }
    this.generateDates();
    this.colorRange = d3.scale.linear().range(["white", this.heatColor]).domain([0, 1]);
    var fd_dow = new Date(Date.UTC(this.year,0,1)).getDay();
    this.firstSunday = new Date(Date.UTC(this.year, 0, 1 - fd_dow));
  },
  props:['year', 'cellSize']
}
</script>

<style scoped>
.calender-map {
  font: 10px sans-serif;
  shape-rendering: crispEdges;
}
.day {
  stroke: #666;
  text-anchor: end;
}
.month {
  fill: white;
}
.border{
  fill: none;
  stroke: #000;
  stroke-width: 2px;
}
.dow{
  fill: white;
  text-anchor: end;
}
</style>
