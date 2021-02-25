/* eslint-disable no-useless-catch */
import instance from '@/http'

const endpoint = '/api/v1/books/'

export default class BookService {
  static async getAllBooks() {
    try {
      const source = await instance.get(endpoint)
      return source.data
    } catch (error) {
      throw error
    }
  }

  static async getBook(id) {
    try {
      const source = await instance.get(endpoint + id)
      return source.data
    } catch (error) {
      throw error
    }
  }

  static async addBook(params) {
    try {
      const result = await instance.post(endpoint, params)
      return result
    } catch (error) {
      throw error
    }
  }

  static async updateBook(id, params) {
    try {
      await instance.put(endpoint + id, params)
    } catch (error) {
      throw error
    }
  }

  static async deleteBook(id) {
    try {
      await instance.delete(endpoint + id)
    } catch (error) {
      throw error
    }
  }
}
