import {createContext, Dispatch, JSX, ReactNode, SetStateAction, useContext, useEffect, useState} from "react"
import styles from "./styles.module.scss"
import useViewTransition, {viewRef} from "@/hooks/useViewTransition/useViewTransition"


const MenuContext = createContext<string>("")

type Props = {
    children: ReactNode
    width?: number
    open?: boolean
    setOpen?: Dispatch<SetStateAction<boolean>>
    defaultView: string
}

export function Menu({children, width, open, setOpen, defaultView}: Props) {
    let refs: viewRef[] = [] // We can't use state here because it would cause an infinite loop... I definitely did not spend 1.5 hours on this
    const [attachedRefs, setAttachedRefs] = useState<boolean>(false)

    useEffect(() => {
        if (Array.isArray(children)) {
            let allMenuViews = true
            children.forEach((Child) => {
                if (Child.type.name !== "View") {
                    allMenuViews = false
                    return
                }
            })
            if (!allMenuViews) {
                throw new Error("Menu must have at least one View")
            }
        } // TODO: this is not optimal => allow no view but multiple of the other components
    }, [children])
    const [view, setView] = useState<string>(defaultView)
    const [lastView, setLastView] = useState<string>("")


    // if (!open) {
    //     return null
    // } // DEBUG

    function Children() {
        if (!Array.isArray(children)) return <></>
        let refs: viewRef[] = []

        let childArray = children.map((child, index) => {
            const [ref, setRef] = useState<HTMLDivElement | null>(null)
            useEffect(() => {
                if (ref) {
                    refs.push({element: ref, id: child.props.id})
                }
            }, [ref])


            return (
                <div key={index} ref={setRef} style={
                    {
                        minWidth: width ? width : "100%",
                    }
                }>
                    {child}
                </div>
            )
        })

        const r = useViewTransition(view, lastView, refs)
        return (
            <div ref={r}>
                {childArray as unknown as JSX.Element}
            </div>
        )
    }

    function change() {
        if (view === "default") {
            setLastView("default")
            setView("anotherView")
        }
        if (view === "anotherView") {
            setLastView("anotherView")
            setView("default")
        }
    }

    return (
        <MenuContext.Provider value={defaultView}>
            <div className={styles.menu}>
                <Children/>
            </div>
            <button onClick={change}/>
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
        // hmm, how do we do this? - We do it with a Context!
        return (
            <div className={styles.view}>
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