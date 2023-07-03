import {ReactNode} from "react"

export function Menu() {
    return (
        <div>
            Menu
        </div>
    )
}


export namespace Menu {
    type ViewProps = {
        children: ReactNode
        id: string
        name: string
    }

    export function View({children, id, name}: ViewProps) {
        return (
            <div>
                View
            </div>
        )
    }



    type ItemProps = {
        children: ReactNode
        icon: ReactNode
        label: ReactNode
        color: string
        toView: string
        onSelect: () => void
    }

    export function Item({children, icon, label, color}: ItemProps) {
        return (
            <div>
                Item
            </div>
        )
    }

    export function Divider() {
        return (
            <div>
                Divider
            </div>
        )
    }

    type LabelProps = {
        children: ReactNode
        icon: ReactNode
        label: ReactNode

    }

    export function Label({children, icon, label}: LabelProps) {
        return (
            <div>
                Label
            </div>
        )
    }
}