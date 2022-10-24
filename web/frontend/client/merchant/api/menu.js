const Menu = [
  {header: 'Index'},
  {
    title: 'Dashboard',
    group: 'dashboard',
    icon: 'dashboard',
    name: 'Dashboard',
    href: '/'
  },
  {header: 'Reservation'},
  {
    title: 'History',
    group: 'apps',
    icon: 'book_online',
    name: 'history',
    href: '/history'
  },
  {header: 'Merchant'},
  {
    title: 'edit',
    group: 'merchant',
    component: 'components',
    icon: 'store',
    name: 'merchant',
    href: '/merchant/edit'
  },
  {divider: true},
  {
    title: 'Logout',
    group: 'extra',
    icon: 'logout',
    href: '/logout'
  },
];
// reorder menu
Menu.forEach((item) => {
  if (item.items) {
    item.items.sort((x, y) => {
      let textA = x.title.toUpperCase();
      let textB = y.title.toUpperCase();
      return (textA < textB) ? -1 : (textA > textB) ? 1 : 0;
    });
  }
});

export default Menu;
