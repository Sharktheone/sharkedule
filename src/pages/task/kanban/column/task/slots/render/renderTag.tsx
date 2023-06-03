import {TagsSlot} from "@kanban/column/task/slots/slotTypes"
import {kanbanTagType} from "@kanban/types"



export default function RenderTags(tagSlot: TagsSlot) {
    let tag = tagSlot.tag
    return (
        <div>
            {tag.map((tag: kanbanTagType) => (
                <RenderTag tag={tag}/>
            ))}
        </div>
    )
}

type RenderTagProps = {
    tag: kanbanTagType
}

export function RenderTag({tag}: RenderTagProps) {
    return (
        <div>
            {tag.name}
        </div>
    )
}