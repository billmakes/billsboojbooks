import NavHeader from '@/components/nav/NavHeader'

const fullLayout = {
  'nav-header': NavHeader
}

export default [
  {
    path: '/',
    name: 'my-books',
    label: 'My Books',
    components: {
      default: () => import('@/components/views/MyBooks'),
      ...fullLayout
    }
  },
  {
    path: '/add-book/',
    name: 'add-book',
    components: {
      default: () => import('@/components/views/AddBook'),
      ...fullLayout
    }
  },
  {
    path: '/book/:id',
    name: 'book',
    props: {
      default: true
    },
    components: {
      default: () => import('@/components/views/BookDetails'),
      ...fullLayout
    }
  },

  {
    path: '*',
    name: 'wild',
    component: () =>
      import(/* webpackChunkName: "error" */ '@/components/views/Error'),
    props: {
      code: 'notFound'
    }
  }
]
