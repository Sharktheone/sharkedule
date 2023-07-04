import {Dispatch, ReactNode, SetStateAction} from "react"
import styles from "./menu.module.css"


type Props = {
    children: ReactNode
    width?: number
    open?: boolean
    setOpen?: Dispatch<SetStateAction<boolean>>

}
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
        // hmm, how do we do this?
        return (
            <div>
                View
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