<template>
  <div>
    <div class="mt-6 mb-5">
    <Search @search-submitted="getResults"></Search>
  </div>
  <div v-if="results.length === 0 && ussed">
    <AlertToast @close-toast="closeToast" message="No results found"></AlertToast>
  </div>
  <div v-if="results.length > 0">
    <div v-for="result in results" :key="result._id">
      <Result
        :from="result._source.From"
        :to="result._source.To"
        :date="result._source.Date"
        :subject="result._source.Subject"
        :body="result._source.bodyData"
      />
    </div>
  </div>
  </div>
</template>

<script>
import Result from './components/Result.vue'
import Search from './components/Search.vue'
import AlertToast from './components/AlertToast.vue'
import axios from 'axios'

export default {
  name: 'App',
  data() {
    return {
      results: [],
      ussed: false
    }
  },
  methods:{
    async getResults(searchTerm) {
      console.log(searchTerm);
      const response = await axios.post(
        process.env.VUE_APP_API_SEARCH_URL,
        {
          search_type: "match",
          query: {
            term: searchTerm,
            field: "_all"
          },
          sort_fields: ["-@timestamp"],
          from: 0,
          max_results: 5,
          _source: [
            "Subject", "From", "bodyData", "To", "Date"
          ]
        }
      ).then(response => {
        console.log(response);
        if ('Hits' in response.data === true && 'hits' in response.data.Hits === true) {
          response.data.Hits.hits.forEach(element => {
            try{

              element._source.bodyData = decodeURIComponent(element._source.bodyData)
            } catch (error) {
              console.warn(element._source.bodyData)
              console.warn(error);
            }
          });
        }
        return response.data;
      }).then(data => {
        if ('Hits' in data === true || 'hits' in data.Hits === true) {
          data.Hits.hits.sort((a, b) => {
            return b._score - a._score;
          });
        }
        return data;
      })
      .catch(error => {
        console.log(error);
      });
      if ("Hits" in response === false || "hits" in response.Hits === false) {
        this.results = [];
        this.ussed = true;
        return;
      }
      this.results = response.Hits.hits;
      this.ussed = true;
    },
    closeToast () {
      this.ussed = false;
    }
  },
  components: {
    Result,
    Search,
    AlertToast
  }
}
</script>

<style scoped>
article {
  background-color: rgba(255, 255, 255, 0.7); /* Blanco con 50% de transparencia */
}
</style>