import Details from "@/components/details/details"

type Props = {
    open: boolean
}

export default function TaskDetails({open}: Props) {
    function onClose() {

    }
    return (
        <Details open={open} onClose={onClose}>
            <div>
                hello
            </div>
        </Details>
    )

}