<template>
  <div class="email-visualizer">
    <nav class="navbar">
      <div class="logo">
        <span class="brand">Email Visualizer</span>
      </div>
      <div class="separator"></div> 
      <div class="search-form">
        <form @submit.prevent="search">
          <input v-model="searchVal" type="text" class="search-input" placeholder="Search">
        </form>
      </div>
    </nav>

    <main class="main-content">
      <div class="table-container">
        <List :fields="fields" :data="searchResults" returnField="Message" @row-clicked="rowClicked"/>
      </div>
      <textarea v-model="textAreaVal" class="body-display" placeholder="Selected email body" readonly></textarea>
    </main>
  </div>
</template>

<script>
import List from "@/components/List.vue";

export default {
  name: 'EmailVisualizer',
  components: { List },
  data() {
    return {
      searchVal: '',
      searchResults: [],
      fields: ["Subject", "From", "To"],
      textAreaVal: 'Message',
    };
  },
  methods: {
    async search() {
      console.log(`Searching ${this.searchTerm}`);

      try {
        const response = await fetch("http://localhost:4040/api/search", {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            search_term: this.searchTerm,
            page: 1
          })
        });

        if (!response.ok) {
          throw new Error('Network response was not ok');
        }

        const data = await response.json();
        this.searchResults = data.map(x => x._source);
        console.log(this.searchResults);
      } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
      }
    },
    rowClicked(body) {
      this.textAreaVal = body;
    }
  }
}
</script>
