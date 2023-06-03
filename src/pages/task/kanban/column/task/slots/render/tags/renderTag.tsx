import {TagsSlot} from "@kanban/column/task/slots/slotTypes"
import {kanbanTagType} from "@kanban/types"


type Props = {
    tagSlot: TagsSlot
}

export default function RenderTags({tagSlot}: Props) {
    let tags = tagSlot.tag
    return (
        <div>
            {tags.map((tag: kanbanTagType) => (
                <RenderTag key={tag.uuid} tag={tag}/>
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