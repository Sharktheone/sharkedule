import {Modal, TextInput} from "@mantine/core"
import {useDisclosure} from "@mantine/hooks"

type props = {
    close: () => void,
    opened: boolean
}


export default function CreateNewModal({opened, close}: props) {

    return (
        <Modal opened={opened} onClose={close} title="Create new Kanban Board">
            <form>
                <TextInput withAsterisk label="Name" placeholder="Board Name" required/>
            </form>

        </Modal>
    )
}