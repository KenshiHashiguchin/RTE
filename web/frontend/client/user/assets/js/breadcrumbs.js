const routes = [
  {
    key: 'dashboard',
    parent: null,
    text: 'Dashboard',
    disabled: false,
    href: '/dashboard',
  },
  {
    key: 'merchant',
    parent: 'dashboard',
    text: 'Merchant',
    disabled: false,
    href: '/merchant',
  },
  {
    key: 'reserve',
    parent: 'dashboard',
    text: 'Reservations',
    disabled: false,
    href: '/reserve',
  },
]

function get(key){
  let breadcrumbs = []
  let isLoop = true
  let searchKey = key
  let disabled = true
  while (isLoop) {
    const item = routes.find(v => v.key === searchKey);
    if(item.parent === null){
      breadcrumbs.unshift(item)
      isLoop = false
      break
    }
    item.disabled = disabled
    breadcrumbs.unshift(item)
    searchKey = item.parent
    disabled = false
  }
  return breadcrumbs
}

export default {
  get
}
