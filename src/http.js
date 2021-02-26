import Axios from 'axios'

const instance = Axios.create({
  baseURL: 'https://booklistbooj.herokuapp.com/'
})

export default instance
