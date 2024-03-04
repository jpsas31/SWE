<template>
  <table id="list" class="w-full overflow-x-auto bg-[#333333] text-white shadow-lg">
    <thead>
        <th v-for="field in fields" :key="field" class="bg-[#4a90e2] border text-left">
          {{ field }} 
        </th>
    </thead>
    <tbody>
      <tr v-for="item in data" :key="item.id" @click="$emit('row-clicked', item[returnField])" class="bg-gray-700 hover:bg-[#333333]">
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
