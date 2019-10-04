<template>
  <v-container fluid >
    <v-card v-if="CalendarData.ptotal + CalendarData.utotal > 0">
     <v-card-text >
      <!-- <div class="text-xs-center"> -->
        <h2 class="text-xs-center">{{CalendarData.title}}</h2>
        <!-- </div> -->
      </v-card-text>
        <v-layout row>
          <v-flex class="calender-map">
            <svg viewBox="0 0 700 135" xmlns="http://www.w3.org/2000/svg" v-if="CalendarData.year">
              <!-- <text transform="translate(350,10)" font-size="10px"  font="sans-serif" width="100%" style="text-anchor: middle;" fill="black">{{CalendarData.title}}</text> -->
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
              :x="getX(d)" :y="getY(d)" :fill="getColor(d)" :fill-opacity="getOpacity(d)">
              <title v-if="isLeaveDay(d)">{{getHoverText(d)}}</title>
            </rect>

            <path v-for="m in 12" :d="getPath(m)"
            class="border" fill="none" stroke="#000" stroke-width="2px"></path>
          </g>
        </svg>
      </v-flex>
    </v-layout>
   
  <v-card-actions>
    <v-flex xs2 offset-xs9 >
      <v-radio-group v-model="viewMode" column>
        <v-radio label="All" value="total" color="blue darken-3"></v-radio>
        <v-radio label="Planned" value="planned" color="green darken-3"></v-radio>
        <v-radio label="Unplanned" value="unplanned" color="red darken-2"></v-radio>
      </v-radio-group>
    </v-flex>
  </v-card-actions>
  </v-card>
</v-flex> 
</v-layout>
</v-card-text>
</v-card>
</v-container>
</template>

<script>
import moment from "moment";
import numeral from "numeral";

export default {
  data: () => ({
    showPlanned: true,
    showUnplanned: true,
    legend: {},
    plannedDays:{},
    unplannedDays:{},
    viewMode: "unplanned", // "planned","unplanned","total" 
    AllDays: [],

    cSize :12,
    firstSunday: null,
    dayNames:['Sun','Mon','Tue','Wed','Thu','Fri','Sat'],
    monthNames : [ "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec" ],
    colors : { 
      total:'#0D47A1',
      unplanned:'#D32F2F', //'#BA68C8',
      planned:'#1B5E20'}
  }),
  methods: {
    countDays:function(days){
      var dcount = 0;
      for (var i = 0; i< days.length; i++){
        dcount += days[i].count;
      }
      return dcount;
    },
    countAll: function(upls , pls){
      return this.countDays(upls) + this.countDays(pls);
    },
    getLegendx: function(index){
      return 10*(index)+5;
    },
    getLegendText: function(index){
      return 10*(index)+ index * 4.3 + 8.4;
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
    isLeaveDay: function(daDate){
      var m = moment(daDate);
      var kw = m.format("DD-MM");
      switch(this.viewMode){
        case "planned":
          return ( kw in this.plannedDays  && this.plannedDays[kw] > 0 );
        case "unplanned":
          return (  kw in  this.unplannedDays  && this.unplannedDays[kw] >0 );
        case "total":
        default:
          return ( kw in this.plannedDays  && this.plannedDays[kw] > 0 ) || (  kw in  this.unplannedDays  && this.unplannedDays[kw] >0 );
      }
      
    },
    getHoverText: function(daDate){
      var m = moment(daDate);
      var kw = m.format("DD-MM");
      var dt = m.format("DD-MMM-YYYY") ;
      var count = 0;
      switch(this.viewMode){
        case "planned":
          count = this.plannedDays[kw];
          break;
        case "unplanned":
          count = this.unplannedDays[kw];
          break;
        case "total":
        default:
        if (this.unplannedDays[kw]==undefined && this.plannedDays[kw]==undefined) {
          count = 0;
        } else if (this.plannedDays[kw]==undefined){
          count = this.unplannedDays[kw];
        }else if(this.unplannedDays[kw]==undefined){
          count = this.plannedDays[kw];
        }else{
          count = this.plannedDays[kw] + this.unplannedDays[kw];
          break;
        }
          
      }
      return dt+": " + count + " absences." ;
    },
    getColor: function(daDate) {
      var m = moment(daDate);
      var kw = m.format("DD-MM");
      var color = this.colors[this.viewMode];

      if( this.viewMode === 'planned' &&  ( kw in  this.plannedDays ) && this.plannedDays[kw] > 0 ){
        return color;
      }
      if( this.viewMode === 'unplanned' && ( kw in  this.unplannedDays ) && this.unplannedDays[kw] >0 ){
        return color;
      }
      if ( this.viewMode === 'total' && ( kw in  this.plannedDays || kw in  this.unplannedDays ) ){
        return color;
      }
      if ( m.weekday()%6==0){
          return "#595959";
        }
      return "#424242";
    },
    getOpacity: function(daDate) {
      var m = moment(daDate);
      var kw = m.format("DD-MM");
      var count = 0;
      var min = 0, max = 0;
      switch(this.viewMode){
        case "planned":
          max = this.CalendarData.pmax;
          min = this.CalendarData.pmin;
          if (kw in this.plannedDays){
            count = this.plannedDays[kw];
          }
          break;
        case "unplanned":
          max = this.CalendarData.umax;
          min = this.CalendarData.umin;
          if (kw in this.unplannedDays){
            count = this.unplannedDays[kw];
          }
          break;
        case "total":
        default:
          max = Math.max(this.CalendarData.pmax,this.CalendarData.umax) ;
          min = Math.min(this.CalendarData.pmin,this.CalendarData.umin) ;
          if (kw in this.plannedDays){
            count = this.plannedDays[kw];
          }
          if (kw in this.unplannedDays){
            count += this.unplannedDays[kw];
          }
      } 

      if (count == 0){
        return 1.0;
      }
      var base = 0.22;
      return (count - min)/(max - min) * (1 - base) + base  ;
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
    this.console = window.console;
    this.generateDates();
    var fd_dow = new Date(Date.UTC(this.CalendarData.year,0,1)).getDay();
    this.firstSunday = new Date(Date.UTC(this.CalendarData.year, 0, 1 - fd_dow));
    this.plannedDays = this.CalendarData.pdays;
    this.unplannedDays = this.CalendarData.udays;
    
  },
  props:['CalendarData', 'LeaveTypes'],
  filters: {
    percent: function (value) {
      return numeral(value).format("0.00%"); 
    },
    capitalize: function (value) {
    if (!value) return ''
    value = value.toString()
    return value.charAt(0).toUpperCase() + value.slice(1)
    }
  }
}
</script>

