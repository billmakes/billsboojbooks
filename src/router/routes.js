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
    path: '/book/:id',
    name: 'book',
    props: true,
    components: {
      default: () => import('@/components/views/BookDetails'),
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
