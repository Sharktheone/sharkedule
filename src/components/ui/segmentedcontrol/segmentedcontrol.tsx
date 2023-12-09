import {Dispatch, HTMLAttributes, SetStateAction} from "react"

type Props = {
    data: { label: string, value: string }[]
    onChange: (value: string) => void | Dispatch<SetStateAction<string>>
    value: string
} & HTMLAttributes<HTMLDivElement>

export function SegmentedControl({data}: Props) {
    return (
        <div/>
    )
}