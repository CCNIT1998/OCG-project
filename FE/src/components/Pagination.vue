<template>
  <nav aria-label="Page navigation example ">
    <ul class="pagination justify-content-center">
      <li class="page-item">
        <a
          class="page-link"
          href="#"
          @click="changePage(pageIndex - 1, $event)"
          :class="{ disabled: pageIndex === 1 }"
          >Previous</a
        >
      </li>

      <li v-for="page in displayPages" :key="page">
        <a
          class="page-link"
          href="#"
          @click="changePage(page, $event)"
          :class="{
            active: pageIndex ? page === pageIndex : page === 1,
          }"
          >{{ page }}</a
        >
      </li>
      <li class="page-item">
        <button
          href="#"
          @click="changePage(pageIndex + 1, $event)"
          class="page-link"
          :class="{ disabled: pageIndex === length }"
        >
          Next
        </button>
      </li>
    </ul>
  </nav>
</template>

<script>
export default {
  name: 'Pagination',

  props: ['length', 'pageSize', 'pageIndex', 'numberOfDisplayPages'],

  emits: ['change'],

  computed: {
    displayPages() {
      let numberOfDisplay = this.numberOfDisplayPages || 5;
      const halfNumberOfDisplayPages = Math.floor(numberOfDisplay / 2);

      let display = [];

      let startDisplay = this.pageIndex - halfNumberOfDisplayPages;
      if (startDisplay < 0) startDisplay = 0;

      numberOfDisplay--;
      let endDisplay = startDisplay + numberOfDisplay;

      if (endDisplay > this.length - 1) {
        endDisplay = this.length - 1;
        startDisplay = endDisplay - numberOfDisplay;
        if (startDisplay < 0) startDisplay = 0;
      }

      for (let i = startDisplay + 1; i <= endDisplay + 1; i++) {
        display.push(i);
      }

      return display;
    },
  },

  methods: {
    changePage(page, event) {
      event.preventDefault();
      window.scrollTo(0, 0);
      this.$emit('change', page);
    },
  },
};
</script>

<style scoped>
.item-pagination {
  user-select: none;
}
.item-pagination.disabled {
  pointer-events: none;
  opacity: 0.5;
}
.active {
  background-color: #666;
  color: white;
}
</style>
