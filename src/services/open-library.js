import instance from '@/http'

const endpoint = 'https://openlibrary.org'

async function getBookDescription(title) {
  let searchResult = await instance.get(
    `${endpoint}/search.json?q=${title}&_facet=false&_spellcheck_count=0&limit=10&fields=key,cover_i,title,author_name,name&mode=everything`
  )

  let resultKey = searchResult.data.docs[0].key + '.json'

  let bookResult = await instance.get(`${endpoint + resultKey}`)

  if (bookResult.data.description.value)
    return bookResult.data.description.value
  else return 'No description available at this time.'
}

export default getBookDescription
