// this file was generated by protoc-gen-gotemplate

package node

import (
	"github.com/pkg/errors"

	"berty.tech/core/pkg/errorcodes"
	"berty.tech/core/pkg/validate/validator"
)

var (
	_ = errors.New
	_ = validator.IsContactKey
	_ = errorcodes.IsSubCode
)

func (m *TestErrorInput) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: Kind - name:"kind" number:1 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"kind"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	return nil
}
func (m *ProtocolsOutput) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: Protocols - name:"protocols" number:1 label:LABEL_REPEATED type:TYPE_STRING json_name:"protocols"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	return nil
}
func (m *AppVersionOutput) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: Version - name:"version" number:1 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"version"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	return nil
}
func (m *PingDestination) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: Destination - name:"destination" number:1 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"destination"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	return nil
}
func (m *ContactRequestInput) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: ContactID - name:"contact_id" number:1 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"contactId" options:<[]:true []:"ContactID" >  (is_contact_key=true, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	if err := validator.IsContactKey(m.GetContactID()); err != nil {
		return err
	}

	// handling field: ContactOverrideDisplayName - name:"contact_override_display_name" number:2 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"contactOverrideDisplayName"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: IntroText - name:"intro_text" number:3 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"introText" options:<[]:256 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=256, skip=false, required=false, min_items=0, max_items=0)
	if len(m.GetIntroText()) > 256 {
		return errors.New("IntroText must not be longer than 256")
	}
	return nil
}
func (m *ContactAcceptRequestInput) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: ContactID - name:"contact_id" number:1 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"contactId" options:<[]:1 []:true []:"ContactID" >  (is_contact_key=true, defined_only=false, min_len=1, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	if len(m.GetContactID()) < 1 {
		return errors.New("ContactID must be longer than 1")
	}
	if err := validator.IsContactKey(m.GetContactID()); err != nil {
		return err
	}
	return nil
}
func (m *ConversationAddMessageInput) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: Conversation - name:"conversation" number:1 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".berty.entity.Conversation" json_name:"conversation" options:<[]:true []:false >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=true, min_items=0, max_items=0)
	if m.GetConversation() == nil {
		return errors.New("Conversation is required")
	}

	if v, ok := interface{}(m.GetConversation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: Conversation")
		}
	}

	// handling field: Message - name:"message" number:2 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".berty.entity.Message" json_name:"message" options:<[]:true []:false >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=true, min_items=0, max_items=0)
	if m.GetMessage() == nil {
		return errors.New("Message is required")
	}

	if v, ok := interface{}(m.GetMessage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: Message")
		}
	}

	return nil
}
func (m *EventStreamInput) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: Filter - name:"filter" number:1 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".berty.entity.Event" json_name:"filter"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetFilter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: Filter")
		}
	}

	return nil
}
func (m *CommitLog) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: Operation - name:"operation" number:1 label:LABEL_OPTIONAL type:TYPE_ENUM type_name:".berty.node.CommitLog.Operation" json_name:"operation" options:<[]:true >  (is_contact_key=false, defined_only=true, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	// WARN: $fieldName: the 'defined_only' validator is not yet implemented

	// handling field: Entity - name:"entity" number:2 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".berty.node.CommitLog.Entity" json_name:"entity" options:<[]:true []:false >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=true, min_items=0, max_items=0)
	if m.GetEntity() == nil {
		return errors.New("Entity is required")
	}

	if v, ok := interface{}(m.GetEntity()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: Entity")
		}
	}

	return nil
}
func (m *EventListInput) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: Filter - name:"filter" number:1 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".berty.entity.Event" json_name:"filter"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetFilter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: Filter")
		}
	}

	// handling field: OnlyWithoutAckedAt - name:"only_without_acked_at" number:2 label:LABEL_OPTIONAL type:TYPE_ENUM type_name:".berty.node.NullableTrueFalse" json_name:"onlyWithoutAckedAt"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: OnlyWithoutSeenAt - name:"only_without_seen_at" number:3 label:LABEL_OPTIONAL type:TYPE_ENUM type_name:".berty.node.NullableTrueFalse" json_name:"onlyWithoutSeenAt"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: Paginate - name:"paginate" number:99 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".berty.node.Pagination" json_name:"paginate"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetPaginate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: Paginate")
		}
	}

	return nil
}
func (m *ContactListInput) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: Filter - name:"filter" number:1 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".berty.entity.Contact" json_name:"filter"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetFilter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: Filter")
		}
	}

	// handling field: Paginate - name:"paginate" number:99 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".berty.node.Pagination" json_name:"paginate"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetPaginate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: Paginate")
		}
	}

	return nil
}
func (m *ContactInput) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: Filter - name:"filter" number:1 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".berty.entity.Contact" json_name:"filter"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetFilter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: Filter")
		}
	}

	return nil
}
func (m *ConversationListInput) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: Filter - name:"filter" number:1 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".berty.entity.Conversation" json_name:"filter"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetFilter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: Filter")
		}
	}

	// handling field: Paginate - name:"paginate" number:99 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".berty.node.Pagination" json_name:"paginate"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetPaginate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: Paginate")
		}
	}

	return nil
}
func (m *ConversationCreateInput) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: Contacts - name:"contacts" number:1 label:LABEL_REPEATED type:TYPE_MESSAGE type_name:".berty.entity.Contact" json_name:"contacts" options:<[]:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=1, max_items=0)

	if v, ok := interface{}(m.GetContacts()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: Contacts")
		}
	}

	if len(m.GetContacts()) < 1 {
		return errors.New("Contacts must contain at least 1 item(s)")
	}

	// handling field: Title - name:"title" number:2 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"title"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: Topic - name:"topic" number:3 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"topic"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: Kind - name:"kind" number:4 label:LABEL_OPTIONAL type:TYPE_ENUM type_name:".berty.entity.Conversation.Kind" json_name:"kind"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	return nil
}
func (m *ConversationManageMembersInput) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: Conversation - name:"conversation" number:1 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".berty.entity.Conversation" json_name:"conversation" options:<[]:true []:false >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=true, min_items=0, max_items=0)
	if m.GetConversation() == nil {
		return errors.New("Conversation is required")
	}

	if v, ok := interface{}(m.GetConversation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: Conversation")
		}
	}

	// handling field: Contacts - name:"contacts" number:2 label:LABEL_REPEATED type:TYPE_MESSAGE type_name:".berty.entity.Contact" json_name:"contacts" options:<[]:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=1, max_items=0)

	if v, ok := interface{}(m.GetContacts()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: Contacts")
		}
	}

	if len(m.GetContacts()) < 1 {
		return errors.New("Contacts must contain at least 1 item(s)")
	}
	return nil
}
func (m *DevicePushConfigCreateInput) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: RelayPubkey - name:"relay_pubkey" number:1 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"relayPubkey" options:<[]:"RelayPubkey" >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: PushID - name:"push_id" number:2 label:LABEL_OPTIONAL type:TYPE_BYTES json_name:"pushId" options:<[]:"PushID" >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: PushType - name:"push_type" number:3 label:LABEL_OPTIONAL type:TYPE_ENUM type_name:".berty.push.DevicePushType" json_name:"pushType"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	return nil
}
func (m *Pagination) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: OrderBy - name:"order_by" number:1 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"orderBy"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: OrderDesc - name:"order_desc" number:2 label:LABEL_OPTIONAL type:TYPE_BOOL json_name:"orderDesc"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: First - name:"first" number:11 label:LABEL_OPTIONAL type:TYPE_INT32 json_name:"first"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: After - name:"after" number:12 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"after"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: Last - name:"last" number:13 label:LABEL_OPTIONAL type:TYPE_INT32 json_name:"last"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: Before - name:"before" number:14 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"before"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	return nil
}
func (m *IntegrationTestInput) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: Name - name:"name" number:1 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"name"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	return nil
}
func (m *IntegrationTestOutput) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: Name - name:"name" number:1 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"name"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: Success - name:"success" number:2 label:LABEL_OPTIONAL type:TYPE_BOOL json_name:"success"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: Verbose - name:"verbose" number:3 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"verbose"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: StartedAt - name:"started_at" number:4 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"startedAt" options:<65001:0 65010:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetStartedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: StartedAt")
		}
	}

	// handling field: FinishedAt - name:"finished_at" number:5 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"finishedAt" options:<65001:0 65010:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetFinishedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: FinishedAt")
		}
	}

	return nil
}
func (m *Void) Validate() error {
	if m == nil {
		return nil
	}

	// handling field:  - name:"T" number:1 label:LABEL_OPTIONAL type:TYPE_BOOL json_name:"T"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	return nil
}
func (m *Bool) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: Ret - name:"ret" number:1 label:LABEL_OPTIONAL type:TYPE_BOOL json_name:"ret"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	return nil
}
func (m *EventIDInput) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: EventID - name:"event_id" number:1 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"eventId" options:<[]:"EventID" >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	return nil
}
func (m *LogStreamInput) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: Continues - name:"continues" number:1 label:LABEL_OPTIONAL type:TYPE_BOOL json_name:"continues"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: LogLevel - name:"log_level" number:2 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"logLevel"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: Namespaces - name:"namespaces" number:3 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"namespaces"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: Last - name:"last" number:4 label:LABEL_OPTIONAL type:TYPE_INT32 json_name:"last"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	return nil
}
func (m *LogEntry) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: Line - name:"line" number:1 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"line"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	return nil
}
func (m *LogfileEntry) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: Path - name:"path" number:1 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"path"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: Filesize - name:"filesize" number:2 label:LABEL_OPTIONAL type:TYPE_INT32 json_name:"filesize"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: CreatedAt - name:"created_at" number:3 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"createdAt" options:<65001:1 65010:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: CreatedAt")
		}
	}

	// handling field: UpdatedAt - name:"updated_at" number:4 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"updatedAt" options:<65001:1 65010:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: UpdatedAt")
		}
	}

	return nil
}
func (m *LogfileReadInput) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: Path - name:"path" number:1 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"path"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	return nil
}
func (m *NodeEvent) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: Kind - name:"kind" number:1 label:LABEL_OPTIONAL type:TYPE_ENUM type_name:".berty.node.Kind" json_name:"kind"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: Attributes - name:"attributes" number:2 label:LABEL_OPTIONAL type:TYPE_BYTES json_name:"attributes"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	return nil
}
