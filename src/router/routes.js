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
    path: '/add-book',
    name: 'add-book',
    label: 'Add New Book',
    components: {
      default: () => import('@/components/views/AddBook'),
      ...fullLayout
    }
  }

  //   {
  //     path: '*',
  //     name: 'wild',
  //     component: () =>
  //       import(/* webpackChunkName: "error" */ '@/router/views/Error'),
  //     props: {
  //       code: 'notFound'
  //     }
  //   }
]
