import {createContext, Dispatch, ReactNode, SetStateAction, useContext} from "react"
import styles from "./styles.module.scss"


const MenuContext = createContext<string>("")

type Props = {
    children: ReactNode
    width?: number
    open?: boolean
    setOpen?: Dispatch<SetStateAction<boolean>>
    defaultView: string
}

export function Menu({children, width, open, setOpen, defaultView}: Props) {
    if (Array.isArray(children)) {
        let allMenuViews = true
        children.forEach((child) => {
            if (child.type.name !== "View") {
                allMenuViews = false
            }
        })
        if (!allMenuViews) {
            throw new Error("Menu must have at least one View")
        }
    } // TODO: this is not optimal => allow no view but multiple of the other components

    // if (!open) {
    //     return null
    // } // DEBUG


    return (
        <MenuContext.Provider value={defaultView}>
            {children}
        </MenuContext.Provider>
    )
}


export namespace Menu {
    type ViewProps = {
        children: ReactNode
        id: string
        name: string
    }

    export function View({children, id, name}: ViewProps) {
        const view = useContext(MenuContext)
        if (view !== id) {
            return null
        }

        // hmm, how do we do this? - We do it with a Context!
        return (
            <div>
                {children}
            </div>
        )
    }


    type ComponentStructureProps = {
        children: ReactNode
        icon?: ReactNode
        label?: ReactNode
        color?: string
        className?: string
    }

    function ComponentStructure({children, icon, label, color, className}: ComponentStructureProps) {
        function getColor() {
            return {
                backgroundColor: color
            }
        }

        return (
            <div className={`${styles.component} ${className}`}>
                {icon ?
                    <div className={styles.icon}>
                        {icon}
                    </div>
                    : null
                }
                {children}
                {label ?
                    <div className={styles.label}>
                        {label}
                    </div>
                    : null
                }
            </div>
        )
    }

    type ItemProps = {
        children: ReactNode
        icon?: ReactNode
        label?: ReactNode
        color?: string
        toView?: string
        onSelect: () => void
    }

    export function Item({children, icon, label, color}: ItemProps) {
        return (
            <ComponentStructure className={styles.item}
                                icon={icon} label={label} color={color}>
                <div className={styles.item}>
                    {children}
                </div>
            </ComponentStructure>
        )
    }

    export function Divider() {
        return (
            <div className={styles.divider}/>
        )
    }

    type SectionProps = {
        children: ReactNode
        icon: ReactNode
        label: ReactNode

    }

    export function Section({children, icon, label}: SectionProps) {
        return (
            <ComponentStructure className={styles.labelComponent}
                                icon={icon} label={label}>
                <div className={styles.labelName}>
                    {children}
                </div>
            </ComponentStructure>
        )
    }
}