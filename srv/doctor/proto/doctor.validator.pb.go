// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: doctor/proto/doctor.proto

package doctorpb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *TimePeriod) Validate() error {
	if this.Start != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Start); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Start", err)
		}
	}
	if this.End != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.End); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("End", err)
		}
	}
	return nil
}
func (this *Schedule) Validate() error {
	for _, item := range this.Monday {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Monday", err)
			}
		}
	}
	for _, item := range this.Tuesday {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Tuesday", err)
			}
		}
	}
	for _, item := range this.Wednesday {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Wednesday", err)
			}
		}
	}
	for _, item := range this.Thursday {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Thursday", err)
			}
		}
	}
	for _, item := range this.Friday {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Friday", err)
			}
		}
	}
	for _, item := range this.Saturday {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Saturday", err)
			}
		}
	}
	for _, item := range this.Sunday {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Sunday", err)
			}
		}
	}
	return nil
}
func (this *Doctor) Validate() error {
	if this.Schedule != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Schedule); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Schedule", err)
		}
	}
	return nil
}
func (this *Medicine) Validate() error {
	return nil
}
func (this *AddPatientRequest) Validate() error {
	if this.PatientId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PatientId", fmt.Errorf(`value '%v' must not be an empty string`, this.PatientId))
	}
	return nil
}
func (this *AddPatientResponse) Validate() error {
	if this.Doctor != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Doctor); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Doctor", err)
		}
	}
	return nil
}
func (this *GetDoctorRequest) Validate() error {
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	return nil
}
func (this *GetDoctorResponse) Validate() error {
	if this.Doctor != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Doctor); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Doctor", err)
		}
	}
	return nil
}
func (this *CreateDoctorRequest) Validate() error {
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if this.Office == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Office", fmt.Errorf(`value '%v' must not be an empty string`, this.Office))
	}
	if nil == this.Schedule {
		return github_com_mwitkow_go_proto_validators.FieldError("Schedule", fmt.Errorf("message must exist"))
	}
	if this.Schedule != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Schedule); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Schedule", err)
		}
	}
	return nil
}
func (this *CreateDoctorResponse) Validate() error {
	if this.Doctor != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Doctor); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Doctor", err)
		}
	}
	return nil
}
func (this *CheckupRequest) Validate() error {
	if this.PatientId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PatientId", fmt.Errorf(`value '%v' must not be an empty string`, this.PatientId))
	}
	if this.DoctorId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DoctorId", fmt.Errorf(`value '%v' must not be an empty string`, this.DoctorId))
	}
	return nil
}
func (this *CheckupResponse) Validate() error {
	return nil
}
