<template>
  <div>
    <b-card class="Book" v-if="book">
      <div>
        <h3>{{ title }}</h3>
        <h5>{{ author }} - {{ year }}</h5>
        <div v-if="bookRead"><BookRead :source="params" /></div>
      </div>
      <BookTags :tags="tags" />
      <div v-if="description">
        <p>{{ description }}</p>
      </div>
      <div v-else>
        <b-spinner label="Loading..."></b-spinner>
      </div>
      <div class="mt-2">
        <b-button v-b-modal.edit-modal class="mr-1">Edit Info</b-button>
        <b-button
          v-b-modal.delete-modal
          class="Book__Options--Button"
          variant="danger"
          >Delete Record</b-button
        >
      </div>
    </b-card>
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
import getBookDescription from '@/services/open-library'
import BookTags from '@/components/books/shared/BookTags'
import BookRead from '@/components/books/shared/BookRead'
const bookFields = ['title', 'author', 'year', 'tags', 'read', 'readDate']

export default {
  name: 'Book',
  components: {
    BookTags,
    BookRead
  },
  data() {
    return {
      book: null,
      title: '',
      author: '',
      year: '',
      tags: [],
      read: null,
      description: null
    }
  },
  props: {
    source: {
      type: Object
    }
  },
  mounted() {
    if (this.source) {
      this.book = this.source
      this.assignFields(this, this.source)
      this.read = this.source.read || false
      getBookDescription(this.source.title)
        .then(desc => (this.description = desc))
        .catch(
          () => (this.description = 'No description available at this time.')
        )
    } else {
      BookService.getBook(this.$route.params.id)
        .then(res => {
          this.book = res.book
          this.assignFields(this, res.book)
          this.read = res.book.read || false
        })
        .catch(() => {
          this.$router.push({ name: 'wild' })
        })
    }
  },
  computed: {
    bookRead() {
      return this.book.read
    },
    params() {
      return {
        title: this.title,
        author: this.author,
        year: this.year,
        tags: this.tags,
        read: this.read || false,
        readDate: this.readDate
      }
    }
  },
  methods: {
    assignFields(assignObj, assignee) {
      bookFields.forEach(field => {
        assignObj[field] = assignee[field]
      })
    },
    cancel() {
      this.$root.$emit('bv::hide::modal', 'edit-modal')
      this.$router.push({ path: '/' })
      this.assignFields(this, this.book)
    },
    save() {
      this.$root.$emit('bv::hide::modal', 'edit-modal')
      BookService.updateBook(this.$route.params.id, this.params).then(() => {
        this.assignFields(this, this.params)
        this.$router.push({ path: '/' })
      })
    },
    deleteBook() {
      this.$root.$emit('bv::hide::modal', 'delete-modal')
      let id = this.$route.params.id || this.source.id
      BookService.deleteBook(id).then(() => {
        this.$router.push({ path: '/' })
      })
    }
  }
}
</script>
