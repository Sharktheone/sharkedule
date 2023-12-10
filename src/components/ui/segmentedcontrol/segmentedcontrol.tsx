import {HTMLAttributes} from "react"

type Props = {
    data: { label: string, value: string }[]
    onChange?:  (value: string) => void
    value: string
    classNames?: CSSModuleClasses
} & Omit<HTMLAttributes<HTMLDivElement>, "onChange">

export function SegmentedControl({data}: Props) {
    //TODO

    return (
        <div/>
    )
}