import {ReactNode} from "react"


export type MenuView = {
    name: string
    id: string
    items: MenuItem[] | ReactNode
}

export type MenuItems = MenuDivider | MenuGroup | MenuItem | MenuLink | MenuToggle | ReactNode

export type MenuDivider = {
    type: "divider"
    name?: string
}

export type MenuGroup = {
    name: string
    icon: ReactNode
    items: MenuItems[]
}

export type MenuItem = {
    name: string
    icon: ReactNode
    onClick: () => void
}

export type MenuLink = {
    name: string
    icon: ReactNode
    to: string
}

export type MenuToggle = {
    name: string
    icon: ReactNode
    view: string
}