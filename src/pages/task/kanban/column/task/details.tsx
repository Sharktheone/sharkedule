import Details from "@/components/details/details"
import {Dispatch, SetStateAction} from "react"

type Props = {
    open: boolean
    setOpen: Dispatch<SetStateAction<boolean>>
}

export default function TaskDetails({open, setOpen}: Props) {
    function onClose() {
        setOpen(false)

    }
    return (
        <Details open={open} onClose={onClose} title="task.name //TODO">
            <div>
                hello
            </div>
        </Details>
    )
}