# Personal TODOs for this project

a bit ironic in a todo app, but hey, it's work in progress

- [ ] Add Some kind of dashboard "root" page
- [ ] Add a nav bar / project shell
- [ ] Settings site for things like details (Drawer/Popup) etc.
- [ ] Activity Log and type
- [x] Think about Progress field => percentage or something else?
- [ ] Dont use contenteditable for the description field etc.
- [ ] Better "closing" of create shit when clicking out
- [ ] Replace Slots?
- [x] Think about custom fields
- [ ] CSS use more classes and remove deeply nested structures
- [x] Remove colors.ts and switch to "inline" mantine colors from colors.scss
- [ ] Only post one notification when adding multiple items
- [x] Think about an `Element` interface. And tasks are only `Elemment`s. But also Boards, Meetings, Subtasks
  are `Element`s
- [ ] Everything is an Element, that contains fields (name, tags, etc. = field)
- [ ] Websocket API?
- [ ] Don't delete things directly, mark them as deleted for users to restore (30 days maybe)
- [ ] new permission system:
    - option to make element "groups" basically, all elements referenced by this/these element
   maybe also referenced ONLY by this/these elements or elements referenced by ALL of these elements
    - User Roles
    - Role overrides: if the user has this role he is not allowed to ..., even he has permission from role X (higher order)
    - User fields?
    - Custom logic stuff?
    - How to implement this?
       - Who is responsible to check permissions?
       - Get the User in the HTTP API, check if the user has access and give back the element, on an action performed, check this before getting the element?
       - Or get the "Access" struct and give it the element as arg and the rest of the parameters?
       - Or the idea from above and shove in the element, and it will automatically perform actions on it? - What about when we need to perform actions on multiple elements?
  
- [ ] completely rework packages, to fix the "circular dependency" problem for a last time...
- [ ] namespacing for Elements?
- [ ] Ditch Mantine
    - [x] Notifications System => react-toastify
    - [ ] SegmentedControl
    - [x] Button
    - [ ] ColorPicker
    - [x] Drawer
    - [x] Modal
    - [x] useClickOutside
    - [x] Input
    - [x] Textarea
    - [x] Container
    - [x] Title
    - [x] useDisclosure
    - [x] usePrevious
    - [x] useDebouncedState
    - [x] CloseButton
    - ~~[ ] Stack~~
    - [x] Text
    - [ ] Colors => Monokai?
- [ ] radix-ui?