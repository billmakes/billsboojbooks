<template>
  <div>
    <h5>Create Book Record</h5>
    <div>
      <h5>Title</h5>
      <b-input-group class="mb-2">
        <b-input-group-prepend is-text>
          <b-icon icon="tag-fill"></b-icon>
        </b-input-group-prepend>
        <b-form-input
          id="titleInput"
          type="text"
          label="title"
          v-model="title"
          placeholder="Title"
          trim
        ></b-form-input>
      </b-input-group>
      <h5>Author</h5>
      <b-input-group class="mb-2">
        <b-input-group-prepend is-text>
          <b-icon icon="person-fill"></b-icon>
        </b-input-group-prepend>
        <b-form-input
          id="authorInput"
          type="text"
          label="author"
          v-model="author"
          placeholder="Author"
          trim
        ></b-form-input>
      </b-input-group>
      <h5>Year</h5>
      <b-input-group class="mb-2">
        <b-input-group-prepend is-text>
          <b-icon icon="clock-fill"></b-icon>
        </b-input-group-prepend>
        <b-form-input
          id="yearInput"
          type="text"
          label="year"
          v-model="year"
          placeholder="Year"
          trim
        ></b-form-input>
      </b-input-group>
      <div>
        <label for="tags-basic">Type a new tag and press enter</label>
        <b-form-tags input-id="tags-basic" v-model="tags"></b-form-tags>
      </div>
      <div class="mt-2">
        <b-form-checkbox id="readCheckbox" v-model="read" name="readCheckbox">
          Read
        </b-form-checkbox>
      </div>
    </div>
    <div class="mr-auto">
      <b-button @click="cancel()">Cancel</b-button>
      <b-button class="ml-1" variant="primary" @click="save()">Save</b-button>
    </div>
  </div>
</template>
<script>
import { BookStore } from '@/composables/book-provider.js'

export default {
  name: 'AddBook',
  data() {
    return {
      title: '',
      author: '',
      year: '',
      tags: [],
      read: null
    }
  },
  computed: {
    params() {
      return {
        title: this.title,
        author: this.author,
        year: this.year,
        tags: this.tags,
        read: this.read || false
      }
    }
  },
  methods: {
    cancel() {
      this.$router.push({ path: '/' })
    },
    save() {
      BookStore.addBook(this.params)
      this.$router.push({ path: '/' })
    }
  }
}
</script>