<template>
  <v-app id="inspire">
    <v-navigation-drawer
    clipped
    fixed
    v-model="drawer"
    app
    >
    <v-list two-line class="grey darken-2">
      <v-subheader class="grey darken-3 white--text">Direct Reports</v-subheader>
      <template v-for="(emp,index) in emps">
        
      <v-list-tile avatar ripple class="grey darken-1" @click="GetEmployee(emp.ID)" >
        <v-list-tile-content>
          <v-list-tile-title class="indigo--text text--darken-4" v-html="emp.FullName"></v-list-tile-title>
          <v-list-tile-sub-title class="yellow--text" v-html="emp.JobTitle"></v-list-tile-sub-title>
          <v-list-tile-sub-title class="white--text" v-html="emp.Email"></v-list-tile-sub-title>
        </v-list-tile-content>
      </v-list-tile>
      <v-divider v-if="index + 1 < emps.length"></v-divider>
      </template>
    </v-list>
  </v-navigation-drawer>
  <v-toolbar app fixed clipped-left>
    <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
    <img src="static/imgs/4lc.svg" style="width:2%;">
    <v-toolbar-title>Forest</v-toolbar-title>
  </v-toolbar>
  <v-content>
    <v-container fluid>
      <v-layout justify-center align-center>
        <v-flex xs12>
          <!-- dont-fill-mask-blanks -->
          <!-- label="Search Employees" -->
          <v-select
          :filter="customFilter"
          placeholder="First or Family name"
          autocomplete
          :async-loading="loading"
          :items="items"
          solo
          :search-input.sync="search"
          v-model="select"
          color="light-green accent-3"
          dense
          @input="OnChange"
          :rules="[() => select.length > 0 || 'You must choose at least one']"
          ></v-select>
          <!-- <vheatmap v-if="yearcounts.year!=null" -->
            <!-- :CalendarData="yearcounts" -->
            <!-- > </vheatmap>  -->
            <!-- {{m.title}} -->
            <vheatmap v-for="m in yearsLeave" :key="m.id" v-bind:data="m" :CalendarData="m">        </vheatmap>

          </v-flex>
        </v-layout>
      </v-container>
    </v-content>
    <v-footer app fixed>
      <span>&copy; ATOP Workforce Analytics</span>
    </v-footer>
  </v-app>
</template>

<script>
import axios from 'axios';
import vheatmap from './components/vheatmap.vue';
export default {
  data: () => ({
    drawer: false,
    search: null,
    items: [],
    select: [],
    loading: false,
    yearcounts : {days: {}, year: null, title:"Not Data Available"},
    yearsLeave : [],
    emps : [],
    customFilter (item, queryText, itemText) {
      return true;
    }
  }),
  watch: {
    search (val) {
      val && this.querySelections(val)
    }
  },
  methods: {
    getReports(id){
      this.emps = [];
      axios.get("api/emp/"+ id)
        .then((response)  =>  {
            this.search = "";
            this.emps = response.data;
            if (this.emps.length == 0){
              this.drawer = false;
            }else{
              this.drawer = true;
            }
            console.log(this.emps);
          }, (error)  =>  {
            console.log(error);
          });
    },
    getLeaves(id){
      this.loading = true;
      this.yearsLeave = [];
      axios.get("api/list/"+ id + "/10")
        .then((response)  =>  {
          this.items =[];
          this.search = "";
          this.loading = false;
          this.yearsLeave = response.data;
        }, (error)  =>  {
             console.log(error);
        });
    },
    GetEmployee(id){
        this.getReports(id);
        this.getLeaves(id);
    },
    OnChange(){
      var id = this.select;
      this.select = "";
      // this.loading = true;
      this.search = "";
      // this.yearsLeave = [];
      this.GetEmployee(id);
      
    },
    querySelections (v) {
      if (!v || v.length < 3 ){
        this.items = [];
        return;
      }
      axios.get("api/search?query="+v)
      .then((response)  =>  {
        this.items = response.data;
      }, (error)  =>  {
        console.log(error);
      })

    }
  },
  components: {
    'vheatmap': vheatmap
  },
}
</script>
