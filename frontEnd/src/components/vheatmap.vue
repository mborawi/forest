<template>
  <v-container fluid>
    <v-card>
     <v-card-text >
      <!-- <div class="text-xs-center"> -->
        <h2 class="text-xs-center">{{CalendarData.title}}</h2>
        <!-- </div> -->
      </v-card-text>
      <v-card-media>
        <v-layout row v-if="!tableView">
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

            <text x="1.6em" dy="13.3em" fill="black" font-size="7px" font-weight="bold" v-if="showPlanned && CalendarData.pcounts.length > 0">Planned:</text>

            <text x="1.6em" dy="14.8em" fill="black" font-size="7px" font-weight="bold" v-if="showUnplanned && CalendarData.ucounts.length > 0">Unplanned:</text>

            <g v-for="(pl, ind) in CalendarData.pcounts" v-if="showPlanned">
              <rect stroke="#666" :width="cSize/2" :height="cSize/2" 
              :x="getLegendx(ind)+'em'" y="8.7em" :fill="colors[pl.cat_id-1]"> </rect>
              <text :x="getLegendText(ind)+'em'" dy="13.3em" font-size="7px" fill="black">{{pl.cat}}</text>
            </g>
            <g v-for="(ul, ind) in CalendarData.ucounts" v-if="showUnplanned">
              <rect  stroke="#666"  :width="cSize/2" :height="cSize/2" 
              :x="getLegendx(ind)+'em'" y="9.8em" :fill="colors[ul.cat_id-1]"> </rect> 
              <text :x="getLegendText(ind)+'em'" dy="14.8em" fill="black" font-size="7px">{{ul.cat}}</text>
            </g>
          </g>
        </svg>
      </v-flex>
      <v-flex>

      </v-flex>
    </v-layout>
  </v-card-media>
  <v-card-text>
    <v-layout align-top justify-space-around v-if="tableView"   class="tr pop-up-message grey lighten-3">
    <v-flex xs4 class="my-3">
      <v-card color="green lighten-5">
        <v-card-title primary-title>
          <div>
            <h3 class="green--text text--darken-4">Planned Leave</h3>
          </div>
        </v-card-title>
        <v-data-table
        :headers="headers"
        :items="CalendarData.pcounts"
        hide-actions
        class="elevation-1 pa-2"
        >
        <template slot="items" slot-scope="props">
          <!-- <td class="text-xs-right">{{ props.item.name }}</td> -->
          <td class="text-xs-left">{{ props.item.cat }}</td>
          <td class="text-xs-center">{{ props.item.count }}</td>
          <td class="text-xs-center">{{ props.item.count / countDays(CalendarData.pcounts) | percent }}</td>
        </template>
      </v-data-table>
    </v-card>
  </v-flex>
   <v-flex xs3 class="my-3"> 
        <v-card color="indigo lighten-5">
          <v-card-title primary-title>
            <div>
              <h3 class="indigo--text text--darken-4">Unplanned Absences</h3>
            </div>
          </v-card-title>
          <v-data-table
          :headers="dowh"
          :items="CalendarData.dows"
          disable-initial-sort
          hide-actions
          class="elevation-1 pa-2">
          <template slot="items" slot-scope="props">
            <td class="text-xs-left">{{ props.item.day | capitalize }}</td>
            <td class="text-xs-center">{{ props.item.count }}</td>
          </template>
        </v-data-table>
      </v-card>
    </v-flex> 
  <v-flex xs4 class="my-3">
    <v-card color="red lighten-5">
      <v-card-title primary-title>
        <div>
          <h3 class="red--text text--darken-4">Unplanned Leave</h3>
        </div>
      </v-card-title>
      <v-data-table
      :headers="headers"
      :items="CalendarData.ucounts"
      hide-actions
      class="elevation-1 pa-2">
      <template slot="items" slot-scope="props">
        <!-- <td class="text-xs-right">{{ props.item.name }}</td> -->
        <td class="text-xs-left">{{ props.item.cat }}</td>
        <td class="text-xs-center">{{ props.item.count }}</td>
        <td class="text-xs-center">{{ props.item.count / countDays(CalendarData.ucounts) | percent }}</td>
      </template>
    </v-data-table>
  </v-card>
</v-flex> 
</v-layout>
</v-card-text>
<v-card-actions>
  <v-btn color="pink darken-4" flat v-on:click="flipView()" v-if="!tableView" >
    <v-icon left>table_chart</v-icon>
    Statistics
  </v-btn>
  <v-spacer v-if="!tableView"></v-spacer>
  <v-flex xs2 justify-end v-if="!tableView"> 
  <v-switch
      v-if="!tableView"
      v-model="showPlanned"
      label="Planned Leaves"
      color="green darken-4"
      hide-details
    ></v-switch>
    <v-switch
      v-if="!tableView"
      v-model="showUnplanned"
      label="Unplanned Leaves"
      color="red darken-4"
      hide-details
    ></v-switch>
    </v-flex>
  <v-btn color="info" flat v-on:click="flipView()" v-if="tableView" >
    <v-icon left dark>insert_chart_outlined</v-icon>
    Calendar
  </v-btn>
  <v-spacer v-if="tableView"></v-spacer>
  <v-btn color="amber darken-4" flat v-on:click="getCSV(CalendarData.file_title)" v-if="tableView">
    <v-icon left>save_alt</v-icon>
    Export CSV
  </v-btn>
  <!-- <v-btn color="info"   v-on:click="getSvg(CalendarData.file_title)">Export</v-btn> -->

</v-card-actions>
</v-card>
<!-- <v-layout align-end justify-end> -->
<!--   <v-flex >
    <div v-for="c in colors" style="height:1.4em;  width: 1.4em; border-color: #00E676; margin:2px; background:coral;" v-bind:style="{background:c}">
      {{c}}
    </div>
  </v-flex>  -->
  <!-- <v-btn color="red" outline v-on:click="flipView()" v-if="!tableView" >Stats</v-btn> -->
  <!-- <v-btn color="info" outline v-on:click="flipView()" v-if="tableView" >Calendar</v-btn> -->
  
  <!-- </v-layout> -->
</v-container>
</template>

<script>
import * as d3 from "d3";
import moment from "moment";
import numeral from "numeral";

export default {
  data: () => ({
    showPlanned: true,
    showUnplanned: true,
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
    dowh: [
    {text: 'Weekday' , value:'day',   align:'left',  class:'grey lighten-4' },
    {text: 'Absences', value:'count', align:'left', class:'grey lighten-4'}],
    headers:[
    // { text: 'Leave Type', value: 'name' },
    { text: 'Category', value: 'cat', align:'center', class:'grey lighten-4' },
    { text: 'Count', value: 'count', align:'center', class:'grey lighten-4'  },
    { text: 'Percentage', value: '', align:'center', class:'grey lighten-4'  },
    ],
    colors : [
    '#00897B',
    '#1976D2',
    '#29B6F6',
    '#2E7D32',
    '#78909C',
    '#7986CB',
    '#7CB342',
    '#7E57C2',
    '#81C784',
    '#A1887F',
    '#BA68C8',
    '#C6FF00',
    '#D4E157',
    '#E0E0E0',
    '#E6EE9C',
    '#EC407A',
    '#F4511E',
    '#EF5350',
    '#FF6F00',
    '#FFA726'  ]
  }),
  methods: {
    flipView: function(){
      this.tableView = !this.tableView;
    },
    isLeaveDay: function(daDate){
      var m = moment(daDate);
      var kw = m.format("DD-MM");
      if (kw in  this.CalendarData.days){
        return true;
      }
      return false;
    },
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
    getHoverText: function(daDate){
      var m = moment(daDate);
      var dt = m.format("DD-MMM-YYYY") ;
      var kw = m.format("DD-MM");
      return dt+": ( "+this.LeaveTypes[this.CalendarData.days[kw].type_id].name+ " )";
    },
    getLegendx: function(index,planned=true){
      return 10*(index)+5;
    },
    getLegendText: function(index,planned=true){
      return 10*(index)+ index * 4.3 + 8.4;
    },
    getCSV: function(fileName){
      var csv = "Planned Leave,,,Unplanned Leave,,\r\n";
      csv += "Category,Count,Percentage,Category,Count,Percentage\r\n";
      var lncnt = Math.max(this.CalendarData.pcounts.length, this.CalendarData.ucounts.length);
      var plcnt = this.countDays(this.CalendarData.pcounts);
      var unplcnt = this.countDays(this.CalendarData.ucounts);

      for(var i=0; i < lncnt;i++){
        if (!this.CalendarData.pcounts[i]){
          csv += ",,,";
        }else{
          csv += this.CalendarData.pcounts[i].cat + "," + this.CalendarData.pcounts[i].count + ","+   this.CalendarData.pcounts[i].count/plcnt * 100 + "%,";
        }
        if(!this.CalendarData.ucounts[i]){
          csv += ",,,\r\n";
        }else{
          csv += this.CalendarData.ucounts[i].cat + "," + this.CalendarData.ucounts[i].count + ","+ this.CalendarData.ucounts[i].count/unplcnt * 100+ "% \r\n";
        }
      }
      csv +="\r\n\r\nTotal Planned,"+ this.countDays(this.CalendarData.pcounts)+",100%";
      csv +=",Total Unplanned,"+ this.countDays(this.CalendarData.ucounts)+",100%,\r\n";

      console.log(csv);
      const blob = new Blob([csv], {type: 'text/csv'})
      const e = document.createEvent('MouseEvents'),
      a = document.createElement('a');
      a.download = fileName + ".csv";
      a.href = window.URL.createObjectURL(blob);
      a.dataset.downloadurl = ['text/csv', a.download, a.href].join(':');
      e.initEvent('click', true, false, window, 0, 0, 0, 0, 0, false, false, false, false, 0, null);
      a.dispatchEvent(e);
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
      var m = moment(daDate);
      var kw = m.format("DD-MM");
      var z = 0.0;

      if (!( kw in  this.CalendarData.days ) || (!this.showPlanned && this.CalendarData.days[kw].name_id==1) || (!this.showUnplanned && this.CalendarData.days[kw].name_id==2)){
        if ( m.weekday()%6==0){
          return "#595959";
        }
        return "#424242";
      }
      
      return this.colors[this.CalendarData.days[kw].cat_id-1];
    },
    getOpacity: function(daDate) {
      var m = moment(daDate);
      var kw = m.format("DD-MM");
      if ( kw in  this.CalendarData.days &&  m.weekday()%6==0){
        return 0.75;
      }
      return 1.0;

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
    this.generateDates();
    this.colorRange = d3.scale.linear().range(["white", this.heatColor]).domain([0, 1]);
    var fd_dow = new Date(Date.UTC(this.CalendarData.year,0,1)).getDay();
    this.firstSunday = new Date(Date.UTC(this.CalendarData.year, 0, 1 - fd_dow));
    
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

