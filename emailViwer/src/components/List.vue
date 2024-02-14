<template>
  <table id="list" class="table-fixed bg-gray-800 text-white shadow-lg">
    <thead>
      <tr>
        <th v-for="field in fields" :key="field" @click="sortTable(field)" class="bg-blue-600 border text-left">
          {{ field }} 
          <i class="sort-icon" aria-label="Sort Icon"></i>
        </th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="item in data" :key="item.id" @click="$emit('row-clicked', item[returnField])" class="hover:bg-gray-700">
        <td v-for="field in fields" :key="field" class="border">{{ item[field] }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script>
export default {
  name: "List",
  props: {
    fields: {
      type: Array,
      required: true
    },
    data: {
      type: Array,
      required: true
    },
    returnField: {
      type: String,
      required: true
    }
  }
}
</script>
