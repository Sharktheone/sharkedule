import {Slot, SlotNames} from "@kanban/column/task/slots/slotTypes"
import RenderTags from "@kanban/column/task/slots/render/tags/renderTag"


type Props = {
    slot: Slot
}

function Render({slot}: Props) {
    switch (slot.type) {
        case SlotNames.TAGS:
            return <RenderTags tagSlot={slot}/>
        default:
            return null
    }
}