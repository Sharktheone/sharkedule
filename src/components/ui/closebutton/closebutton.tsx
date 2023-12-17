import {Button} from "@/components/ui"
import {HTMLAttributes} from "react"


type Props = {} & HTMLAttributes<HTMLButtonElement>

export function CloseButton({}: Props) {
    return (
        <div>
            <Button>CloseButton</Button>
            <p>TODO: CloseButton</p>
        </div>
    )
}