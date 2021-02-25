<template>
  <div>
    <b-card
      :title="source.title"
      img-src="https://picsum.photos/600/300/?image=25"
      :img-alt="`${source.title} by ${source.author}`"
      img-top
      tag="article"
      style="max-width: 20rem"
      class="m-2"
    >
      <b-card-text>
        Tile
        {{ `${source.title} by ${source.author} (${source.year})` }}
      </b-card-text>
      <div v-if="source.read" class="d-flex justify-content-between">
        <b-button :href="`book/${source.id}`" variant="primary"
          >Details</b-button
        >
        <div v-if="source.read">
          Read on {{ source.readDate }}
          <b-icon icon="check" />
        </div>
      </div>
    </b-card>
  </div>
</template>
<script>
export default {
  name: 'BookTile',
  props: {
    source: {
      type: Object,
      required: true
    }
  }
}
</script>