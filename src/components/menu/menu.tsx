import {ReactNode} from "react"
import styles from "./menu.module.css"

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


    type ItemProps = {
        children: ReactNode
        icon?: ReactNode
        label?: ReactNode
        color?: string
        toView?: string
        onSelect: () => void
    }

    export function Item({children, icon, label, color}: ItemProps) {
        function getColor() {
            return {
                backgroundColor: color
            }
        }


        return (
            <div className={styles.labelComponent}>
                {icon ?
                    <div className={styles.icon}>
                        {icon}
                    </div>
                    : null
                }
                <div className={styles.item}>
                    {children}
                </div>
                {label ?
                    <div className={styles.label}>
                        {label}
                    </div>
                    : null
                }
            </div>
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
            <div className={styles.labelComponent}>
                {icon ?
                    <div className={styles.icon}>
                        {icon}
                    </div>
                    : null
                }
                <div className={styles.labelName}>
                    {children}
                </div>
                {label ?
                    <div className={styles.label}>
                        {label}
                    </div>
                    : null
                }
            </div>
        )
    }
}