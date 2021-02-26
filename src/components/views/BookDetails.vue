<template>
  <div>
    <b-card v-show="!deleted" class="Book">
      <h3>{{ title }}</h3>
      <h5>{{ author }} - {{ year }}</h5>
      <div class="Book__Options">
        <b-button v-b-modal.edit-modal class="Book__Options--Button"
          >Edit Info</b-button
        >
        <b-button
          v-b-modal.delete-modal
          class="Book__Options--Button"
          variant="danger"
          >Delete Record</b-button
        >
      </div>
    </b-card>
    <div v-show="deleted">
      <h1 v-if="error">404</h1>
      <p>This book record either does not exist or it may have been deleted.</p>
      <router-link to="/">Click here to go back home.</router-link>
    </div>
    <b-modal id="edit-modal" title="BootstrapVue">
      <template slot="modal-header">
        <h5>Edit Book Record</h5>
      </template>
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
          <p class="mt-2">Value: {{ tags }}</p>
        </div>
        <b-form-checkbox id="readCheckbox" v-model="read" name="readCheckbox">
          Read
        </b-form-checkbox>
      </div>
      <template slot="modal-footer">
        <b-button @click="cancel()">Cancel</b-button>
        <b-button variant="primary" @click="save()">Save</b-button>
      </template>
    </b-modal>
    <b-modal id="delete-modal">
      <template slot="modal-header">
        <h5>Delete Book Record</h5>
      </template>
      <p>Are you sure you want to delete book {{ title }}'s record?</p>
      <template slot="modal-footer" slot-scope="{ hide }">
        <b-button @click="hide('forget')">Cancel</b-button>
        <b-button variant="danger" @click="deleteBook()">Yes</b-button>
      </template>
    </b-modal>
  </div>
</template>
<script>
import BookService from '@/services/'
export default {
  name: 'Book',
  data() {
    return {
      book: [],
      title: '',
      author: '',
      year: '',
      tags: [],
      read: null,
      deleted: false,
      error: false
    }
  },
  mounted() {
    BookService.getBook(this.$route.params.id)
      .then(res => {
        this.book = res.book
        this.title = res.book.title
        this.author = res.book.author
        this.year = res.book.year
        this.tags = res.book.tags
        this.read = res.book.read || false
      })
      .catch(() => {
        this.deleted = true
        this.error = true
      })
  },
  computed: {
    fullName() {
      return `${this.firstName} ${this.lastName}`
    },
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
      this.$root.$emit('bv::hide::modal', 'edit-modal')
      this.title = this.book.title
      this.author = this.book.author
      this.year = this.book.year
      this.tags = this.book.tags
      this.read = this.book.read
    },
    save() {
      this.$root.$emit('bv::hide::modal', 'edit-modal')
      BookService.updateBook(this.$route.params.id, this.params)
      this.title = this.params.title
      this.author = this.params.author
      this.year = this.params.year
      this.tags = this.params.tags
      this.read = this.params.read
    },
    deleteBook() {
      this.$root.$emit('bv::hide::modal', 'delete-modal')
      this.deleted = true
      BookService.deleteBook(this.$route.params.id)
    }
  }
}
</script>