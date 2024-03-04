<template>
  <div  class="flex flex-col h-screen max-h-screen">
    <nav class="flex items-center justify-between bg-[#333333] text-white p-4">
      <div class="text-white font-bold">
        <span class="brand">Email Visualizer</span>
      </div>
      <div class="mr-5"></div> 
      <div class="flex-grow">
        <form @submit.prevent="search">
          <input v-model="searchVal" type="text" class="bg-[#444] text-white border border-gray-600 rounded-md px-2 py-1 w-full" placeholder="Search">
        </form>
      </div>
    </nav>

    <main class="bg-[#222] grid grid-cols-2 gap-4 flex-1 min-h-0">
      <div class="flex-1 border border-[#666] overflow-y-auto overflow-x-hidden">
        <List :fields="fields" :data="searchResults" returnField="Message" @row-clicked="rowClicked"/>
      </div>
      <textarea v-model="textAreaVal" class="max-h-screen w-full overflow-auto bg-[#444] text-white border border-[#666] p-2" placeholder="Selected email body" readonly></textarea>
    </main>
    <footer class="bg-[#333333] p-4 flex items-center justify-center">
      <button @click="prevPage" :disabled="currentPage <= 1" class="bg-[#444] hover:bg-[#555] w-full text-white bg-transparent border border-white rounded px-2 py-1">Previous</button>
      <span class="text-white px-2">{{ currentPage }}</span>
      <button @click="nextPage"  :disabled="currentPage > this.totalPages" class=" bg-[#444] hover:bg-[#555] w-full text-white bg-transparent border border-white rounded px-2 py-1">Next</button>
    </footer>
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
      currentPage: 1,
      totalPages:0
    };
  },
  methods: {
    async search() {
      console.log(`Searching ${this.searchVal}`);

      try { 
        const response = await fetch("http://localhost:4040/api/search", {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            search_term: this.searchVal,
            page: this.currentPage
          })
        });

        if (!response.ok) {
          throw new Error('Network response was not ok');
        }

        const data = await response.json();
       
        if (data === null) {
          console.warn('there are no results with the given keyword');
          this.searchResults = []
          return;
        }
        
        this.searchResults = data.results.map(x => x._source);
        this.totalPages = data.pages;
        console.log(this.totalPages)
        
      } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
      }
    },
    rowClicked(body) {
      this.textAreaVal = body;
    },
    nextPage() {
      if (this.currentPage < this.totalPages) {
        this.currentPage++;
        this.search();
      }
      
    },
    prevPage() {
      if (this.currentPage > 1) {
        this.currentPage--;
        this.search();
      }
    }
  }
}
</script>
