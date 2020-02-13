// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha2

import (
	unsafe "unsafe"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1alpha1 "k8s.io/component-base/config/v1alpha1"
	v1alpha2 "k8s.io/kube-scheduler/config/v1alpha2"
	config "k8s.io/kubernetes/pkg/scheduler/apis/config"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1alpha2.KubeSchedulerConfiguration)(nil), (*config.KubeSchedulerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeSchedulerConfiguration_To_config_KubeSchedulerConfiguration(a.(*v1alpha2.KubeSchedulerConfiguration), b.(*config.KubeSchedulerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeSchedulerConfiguration)(nil), (*v1alpha2.KubeSchedulerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeSchedulerConfiguration_To_v1alpha2_KubeSchedulerConfiguration(a.(*config.KubeSchedulerConfiguration), b.(*v1alpha2.KubeSchedulerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha2.KubeSchedulerLeaderElectionConfiguration)(nil), (*config.KubeSchedulerLeaderElectionConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeSchedulerLeaderElectionConfiguration_To_config_KubeSchedulerLeaderElectionConfiguration(a.(*v1alpha2.KubeSchedulerLeaderElectionConfiguration), b.(*config.KubeSchedulerLeaderElectionConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeSchedulerLeaderElectionConfiguration)(nil), (*v1alpha2.KubeSchedulerLeaderElectionConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeSchedulerLeaderElectionConfiguration_To_v1alpha2_KubeSchedulerLeaderElectionConfiguration(a.(*config.KubeSchedulerLeaderElectionConfiguration), b.(*v1alpha2.KubeSchedulerLeaderElectionConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha2.Plugin)(nil), (*config.Plugin)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_Plugin_To_config_Plugin(a.(*v1alpha2.Plugin), b.(*config.Plugin), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Plugin)(nil), (*v1alpha2.Plugin)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Plugin_To_v1alpha2_Plugin(a.(*config.Plugin), b.(*v1alpha2.Plugin), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha2.PluginConfig)(nil), (*config.PluginConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_PluginConfig_To_config_PluginConfig(a.(*v1alpha2.PluginConfig), b.(*config.PluginConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.PluginConfig)(nil), (*v1alpha2.PluginConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_PluginConfig_To_v1alpha2_PluginConfig(a.(*config.PluginConfig), b.(*v1alpha2.PluginConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha2.PluginSet)(nil), (*config.PluginSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_PluginSet_To_config_PluginSet(a.(*v1alpha2.PluginSet), b.(*config.PluginSet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.PluginSet)(nil), (*v1alpha2.PluginSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_PluginSet_To_v1alpha2_PluginSet(a.(*config.PluginSet), b.(*v1alpha2.PluginSet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha2.Plugins)(nil), (*config.Plugins)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_Plugins_To_config_Plugins(a.(*v1alpha2.Plugins), b.(*config.Plugins), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Plugins)(nil), (*v1alpha2.Plugins)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Plugins_To_v1alpha2_Plugins(a.(*config.Plugins), b.(*v1alpha2.Plugins), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha2.SchedulerAlgorithmSource)(nil), (*config.SchedulerAlgorithmSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_SchedulerAlgorithmSource_To_config_SchedulerAlgorithmSource(a.(*v1alpha2.SchedulerAlgorithmSource), b.(*config.SchedulerAlgorithmSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.SchedulerAlgorithmSource)(nil), (*v1alpha2.SchedulerAlgorithmSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_SchedulerAlgorithmSource_To_v1alpha2_SchedulerAlgorithmSource(a.(*config.SchedulerAlgorithmSource), b.(*v1alpha2.SchedulerAlgorithmSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha2.SchedulerPolicyConfigMapSource)(nil), (*config.SchedulerPolicyConfigMapSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_SchedulerPolicyConfigMapSource_To_config_SchedulerPolicyConfigMapSource(a.(*v1alpha2.SchedulerPolicyConfigMapSource), b.(*config.SchedulerPolicyConfigMapSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.SchedulerPolicyConfigMapSource)(nil), (*v1alpha2.SchedulerPolicyConfigMapSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_SchedulerPolicyConfigMapSource_To_v1alpha2_SchedulerPolicyConfigMapSource(a.(*config.SchedulerPolicyConfigMapSource), b.(*v1alpha2.SchedulerPolicyConfigMapSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha2.SchedulerPolicyFileSource)(nil), (*config.SchedulerPolicyFileSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_SchedulerPolicyFileSource_To_config_SchedulerPolicyFileSource(a.(*v1alpha2.SchedulerPolicyFileSource), b.(*config.SchedulerPolicyFileSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.SchedulerPolicyFileSource)(nil), (*v1alpha2.SchedulerPolicyFileSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_SchedulerPolicyFileSource_To_v1alpha2_SchedulerPolicyFileSource(a.(*config.SchedulerPolicyFileSource), b.(*v1alpha2.SchedulerPolicyFileSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha2.SchedulerPolicySource)(nil), (*config.SchedulerPolicySource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_SchedulerPolicySource_To_config_SchedulerPolicySource(a.(*v1alpha2.SchedulerPolicySource), b.(*config.SchedulerPolicySource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.SchedulerPolicySource)(nil), (*v1alpha2.SchedulerPolicySource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_SchedulerPolicySource_To_v1alpha2_SchedulerPolicySource(a.(*config.SchedulerPolicySource), b.(*v1alpha2.SchedulerPolicySource), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha2_KubeSchedulerConfiguration_To_config_KubeSchedulerConfiguration(in *v1alpha2.KubeSchedulerConfiguration, out *config.KubeSchedulerConfiguration, s conversion.Scope) error {
	if err := v1.Convert_Pointer_string_To_string(&in.SchedulerName, &out.SchedulerName, s); err != nil {
		return err
	}
	if err := Convert_v1alpha2_SchedulerAlgorithmSource_To_config_SchedulerAlgorithmSource(&in.AlgorithmSource, &out.AlgorithmSource, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_int32_To_int32(&in.HardPodAffinitySymmetricWeight, &out.HardPodAffinitySymmetricWeight, s); err != nil {
		return err
	}
	if err := Convert_v1alpha2_KubeSchedulerLeaderElectionConfiguration_To_config_KubeSchedulerLeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	if err := v1alpha1.Convert_v1alpha1_ClientConnectionConfiguration_To_config_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_string_To_string(&in.HealthzBindAddress, &out.HealthzBindAddress, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_string_To_string(&in.MetricsBindAddress, &out.MetricsBindAddress, s); err != nil {
		return err
	}
	if err := v1alpha1.Convert_v1alpha1_DebuggingConfiguration_To_config_DebuggingConfiguration(&in.DebuggingConfiguration, &out.DebuggingConfiguration, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_bool_To_bool(&in.DisablePreemption, &out.DisablePreemption, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_int32_To_int32(&in.PercentageOfNodesToScore, &out.PercentageOfNodesToScore, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_int64_To_int64(&in.BindTimeoutSeconds, &out.BindTimeoutSeconds, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_int64_To_int64(&in.PodInitialBackoffSeconds, &out.PodInitialBackoffSeconds, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_int64_To_int64(&in.PodMaxBackoffSeconds, &out.PodMaxBackoffSeconds, s); err != nil {
		return err
	}
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = new(config.Plugins)
		if err := Convert_v1alpha2_Plugins_To_config_Plugins(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Plugins = nil
	}
	out.PluginConfig = *(*[]config.PluginConfig)(unsafe.Pointer(&in.PluginConfig))
	return nil
}

// Convert_v1alpha2_KubeSchedulerConfiguration_To_config_KubeSchedulerConfiguration is an autogenerated conversion function.
func Convert_v1alpha2_KubeSchedulerConfiguration_To_config_KubeSchedulerConfiguration(in *v1alpha2.KubeSchedulerConfiguration, out *config.KubeSchedulerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha2_KubeSchedulerConfiguration_To_config_KubeSchedulerConfiguration(in, out, s)
}

func autoConvert_config_KubeSchedulerConfiguration_To_v1alpha2_KubeSchedulerConfiguration(in *config.KubeSchedulerConfiguration, out *v1alpha2.KubeSchedulerConfiguration, s conversion.Scope) error {
	if err := v1.Convert_string_To_Pointer_string(&in.SchedulerName, &out.SchedulerName, s); err != nil {
		return err
	}
	if err := Convert_config_SchedulerAlgorithmSource_To_v1alpha2_SchedulerAlgorithmSource(&in.AlgorithmSource, &out.AlgorithmSource, s); err != nil {
		return err
	}
	if err := v1.Convert_int32_To_Pointer_int32(&in.HardPodAffinitySymmetricWeight, &out.HardPodAffinitySymmetricWeight, s); err != nil {
		return err
	}
	if err := Convert_config_KubeSchedulerLeaderElectionConfiguration_To_v1alpha2_KubeSchedulerLeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	if err := v1alpha1.Convert_config_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	if err := v1.Convert_string_To_Pointer_string(&in.HealthzBindAddress, &out.HealthzBindAddress, s); err != nil {
		return err
	}
	if err := v1.Convert_string_To_Pointer_string(&in.MetricsBindAddress, &out.MetricsBindAddress, s); err != nil {
		return err
	}
	if err := v1alpha1.Convert_config_DebuggingConfiguration_To_v1alpha1_DebuggingConfiguration(&in.DebuggingConfiguration, &out.DebuggingConfiguration, s); err != nil {
		return err
	}
	if err := v1.Convert_bool_To_Pointer_bool(&in.DisablePreemption, &out.DisablePreemption, s); err != nil {
		return err
	}
	if err := v1.Convert_int32_To_Pointer_int32(&in.PercentageOfNodesToScore, &out.PercentageOfNodesToScore, s); err != nil {
		return err
	}
	if err := v1.Convert_int64_To_Pointer_int64(&in.BindTimeoutSeconds, &out.BindTimeoutSeconds, s); err != nil {
		return err
	}
	if err := v1.Convert_int64_To_Pointer_int64(&in.PodInitialBackoffSeconds, &out.PodInitialBackoffSeconds, s); err != nil {
		return err
	}
	if err := v1.Convert_int64_To_Pointer_int64(&in.PodMaxBackoffSeconds, &out.PodMaxBackoffSeconds, s); err != nil {
		return err
	}
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = new(v1alpha2.Plugins)
		if err := Convert_config_Plugins_To_v1alpha2_Plugins(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Plugins = nil
	}
	out.PluginConfig = *(*[]v1alpha2.PluginConfig)(unsafe.Pointer(&in.PluginConfig))
	return nil
}

// Convert_config_KubeSchedulerConfiguration_To_v1alpha2_KubeSchedulerConfiguration is an autogenerated conversion function.
func Convert_config_KubeSchedulerConfiguration_To_v1alpha2_KubeSchedulerConfiguration(in *config.KubeSchedulerConfiguration, out *v1alpha2.KubeSchedulerConfiguration, s conversion.Scope) error {
	return autoConvert_config_KubeSchedulerConfiguration_To_v1alpha2_KubeSchedulerConfiguration(in, out, s)
}

func autoConvert_v1alpha2_KubeSchedulerLeaderElectionConfiguration_To_config_KubeSchedulerLeaderElectionConfiguration(in *v1alpha2.KubeSchedulerLeaderElectionConfiguration, out *config.KubeSchedulerLeaderElectionConfiguration, s conversion.Scope) error {
	if err := v1alpha1.Convert_v1alpha1_LeaderElectionConfiguration_To_config_LeaderElectionConfiguration(&in.LeaderElectionConfiguration, &out.LeaderElectionConfiguration, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_KubeSchedulerLeaderElectionConfiguration_To_config_KubeSchedulerLeaderElectionConfiguration is an autogenerated conversion function.
func Convert_v1alpha2_KubeSchedulerLeaderElectionConfiguration_To_config_KubeSchedulerLeaderElectionConfiguration(in *v1alpha2.KubeSchedulerLeaderElectionConfiguration, out *config.KubeSchedulerLeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha2_KubeSchedulerLeaderElectionConfiguration_To_config_KubeSchedulerLeaderElectionConfiguration(in, out, s)
}

func autoConvert_config_KubeSchedulerLeaderElectionConfiguration_To_v1alpha2_KubeSchedulerLeaderElectionConfiguration(in *config.KubeSchedulerLeaderElectionConfiguration, out *v1alpha2.KubeSchedulerLeaderElectionConfiguration, s conversion.Scope) error {
	if err := v1alpha1.Convert_config_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(&in.LeaderElectionConfiguration, &out.LeaderElectionConfiguration, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_KubeSchedulerLeaderElectionConfiguration_To_v1alpha2_KubeSchedulerLeaderElectionConfiguration is an autogenerated conversion function.
func Convert_config_KubeSchedulerLeaderElectionConfiguration_To_v1alpha2_KubeSchedulerLeaderElectionConfiguration(in *config.KubeSchedulerLeaderElectionConfiguration, out *v1alpha2.KubeSchedulerLeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_config_KubeSchedulerLeaderElectionConfiguration_To_v1alpha2_KubeSchedulerLeaderElectionConfiguration(in, out, s)
}

func autoConvert_v1alpha2_Plugin_To_config_Plugin(in *v1alpha2.Plugin, out *config.Plugin, s conversion.Scope) error {
	out.Name = in.Name
	if err := v1.Convert_Pointer_int32_To_int32(&in.Weight, &out.Weight, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_Plugin_To_config_Plugin is an autogenerated conversion function.
func Convert_v1alpha2_Plugin_To_config_Plugin(in *v1alpha2.Plugin, out *config.Plugin, s conversion.Scope) error {
	return autoConvert_v1alpha2_Plugin_To_config_Plugin(in, out, s)
}

func autoConvert_config_Plugin_To_v1alpha2_Plugin(in *config.Plugin, out *v1alpha2.Plugin, s conversion.Scope) error {
	out.Name = in.Name
	if err := v1.Convert_int32_To_Pointer_int32(&in.Weight, &out.Weight, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_Plugin_To_v1alpha2_Plugin is an autogenerated conversion function.
func Convert_config_Plugin_To_v1alpha2_Plugin(in *config.Plugin, out *v1alpha2.Plugin, s conversion.Scope) error {
	return autoConvert_config_Plugin_To_v1alpha2_Plugin(in, out, s)
}

func autoConvert_v1alpha2_PluginConfig_To_config_PluginConfig(in *v1alpha2.PluginConfig, out *config.PluginConfig, s conversion.Scope) error {
	out.Name = in.Name
	out.Args = in.Args
	return nil
}

// Convert_v1alpha2_PluginConfig_To_config_PluginConfig is an autogenerated conversion function.
func Convert_v1alpha2_PluginConfig_To_config_PluginConfig(in *v1alpha2.PluginConfig, out *config.PluginConfig, s conversion.Scope) error {
	return autoConvert_v1alpha2_PluginConfig_To_config_PluginConfig(in, out, s)
}

func autoConvert_config_PluginConfig_To_v1alpha2_PluginConfig(in *config.PluginConfig, out *v1alpha2.PluginConfig, s conversion.Scope) error {
	out.Name = in.Name
	out.Args = in.Args
	return nil
}

// Convert_config_PluginConfig_To_v1alpha2_PluginConfig is an autogenerated conversion function.
func Convert_config_PluginConfig_To_v1alpha2_PluginConfig(in *config.PluginConfig, out *v1alpha2.PluginConfig, s conversion.Scope) error {
	return autoConvert_config_PluginConfig_To_v1alpha2_PluginConfig(in, out, s)
}

func autoConvert_v1alpha2_PluginSet_To_config_PluginSet(in *v1alpha2.PluginSet, out *config.PluginSet, s conversion.Scope) error {
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = make([]config.Plugin, len(*in))
		for i := range *in {
			if err := Convert_v1alpha2_Plugin_To_config_Plugin(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Enabled = nil
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = make([]config.Plugin, len(*in))
		for i := range *in {
			if err := Convert_v1alpha2_Plugin_To_config_Plugin(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Disabled = nil
	}
	return nil
}

// Convert_v1alpha2_PluginSet_To_config_PluginSet is an autogenerated conversion function.
func Convert_v1alpha2_PluginSet_To_config_PluginSet(in *v1alpha2.PluginSet, out *config.PluginSet, s conversion.Scope) error {
	return autoConvert_v1alpha2_PluginSet_To_config_PluginSet(in, out, s)
}

func autoConvert_config_PluginSet_To_v1alpha2_PluginSet(in *config.PluginSet, out *v1alpha2.PluginSet, s conversion.Scope) error {
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = make([]v1alpha2.Plugin, len(*in))
		for i := range *in {
			if err := Convert_config_Plugin_To_v1alpha2_Plugin(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Enabled = nil
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = make([]v1alpha2.Plugin, len(*in))
		for i := range *in {
			if err := Convert_config_Plugin_To_v1alpha2_Plugin(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Disabled = nil
	}
	return nil
}

// Convert_config_PluginSet_To_v1alpha2_PluginSet is an autogenerated conversion function.
func Convert_config_PluginSet_To_v1alpha2_PluginSet(in *config.PluginSet, out *v1alpha2.PluginSet, s conversion.Scope) error {
	return autoConvert_config_PluginSet_To_v1alpha2_PluginSet(in, out, s)
}

func autoConvert_v1alpha2_Plugins_To_config_Plugins(in *v1alpha2.Plugins, out *config.Plugins, s conversion.Scope) error {
	if in.QueueSort != nil {
		in, out := &in.QueueSort, &out.QueueSort
		*out = new(config.PluginSet)
		if err := Convert_v1alpha2_PluginSet_To_config_PluginSet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.QueueSort = nil
	}
	if in.PreFilter != nil {
		in, out := &in.PreFilter, &out.PreFilter
		*out = new(config.PluginSet)
		if err := Convert_v1alpha2_PluginSet_To_config_PluginSet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.PreFilter = nil
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(config.PluginSet)
		if err := Convert_v1alpha2_PluginSet_To_config_PluginSet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Filter = nil
	}
	if in.PreScore != nil {
		in, out := &in.PreScore, &out.PreScore
		*out = new(config.PluginSet)
		if err := Convert_v1alpha2_PluginSet_To_config_PluginSet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.PreScore = nil
	}
	if in.Score != nil {
		in, out := &in.Score, &out.Score
		*out = new(config.PluginSet)
		if err := Convert_v1alpha2_PluginSet_To_config_PluginSet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Score = nil
	}
	if in.Reserve != nil {
		in, out := &in.Reserve, &out.Reserve
		*out = new(config.PluginSet)
		if err := Convert_v1alpha2_PluginSet_To_config_PluginSet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Reserve = nil
	}
	if in.Permit != nil {
		in, out := &in.Permit, &out.Permit
		*out = new(config.PluginSet)
		if err := Convert_v1alpha2_PluginSet_To_config_PluginSet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Permit = nil
	}
	if in.PreBind != nil {
		in, out := &in.PreBind, &out.PreBind
		*out = new(config.PluginSet)
		if err := Convert_v1alpha2_PluginSet_To_config_PluginSet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.PreBind = nil
	}
	if in.Bind != nil {
		in, out := &in.Bind, &out.Bind
		*out = new(config.PluginSet)
		if err := Convert_v1alpha2_PluginSet_To_config_PluginSet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Bind = nil
	}
	if in.PostBind != nil {
		in, out := &in.PostBind, &out.PostBind
		*out = new(config.PluginSet)
		if err := Convert_v1alpha2_PluginSet_To_config_PluginSet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.PostBind = nil
	}
	if in.Unreserve != nil {
		in, out := &in.Unreserve, &out.Unreserve
		*out = new(config.PluginSet)
		if err := Convert_v1alpha2_PluginSet_To_config_PluginSet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Unreserve = nil
	}
	return nil
}

// Convert_v1alpha2_Plugins_To_config_Plugins is an autogenerated conversion function.
func Convert_v1alpha2_Plugins_To_config_Plugins(in *v1alpha2.Plugins, out *config.Plugins, s conversion.Scope) error {
	return autoConvert_v1alpha2_Plugins_To_config_Plugins(in, out, s)
}

func autoConvert_config_Plugins_To_v1alpha2_Plugins(in *config.Plugins, out *v1alpha2.Plugins, s conversion.Scope) error {
	if in.QueueSort != nil {
		in, out := &in.QueueSort, &out.QueueSort
		*out = new(v1alpha2.PluginSet)
		if err := Convert_config_PluginSet_To_v1alpha2_PluginSet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.QueueSort = nil
	}
	if in.PreFilter != nil {
		in, out := &in.PreFilter, &out.PreFilter
		*out = new(v1alpha2.PluginSet)
		if err := Convert_config_PluginSet_To_v1alpha2_PluginSet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.PreFilter = nil
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(v1alpha2.PluginSet)
		if err := Convert_config_PluginSet_To_v1alpha2_PluginSet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Filter = nil
	}
	if in.PreScore != nil {
		in, out := &in.PreScore, &out.PreScore
		*out = new(v1alpha2.PluginSet)
		if err := Convert_config_PluginSet_To_v1alpha2_PluginSet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.PreScore = nil
	}
	if in.Score != nil {
		in, out := &in.Score, &out.Score
		*out = new(v1alpha2.PluginSet)
		if err := Convert_config_PluginSet_To_v1alpha2_PluginSet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Score = nil
	}
	if in.Reserve != nil {
		in, out := &in.Reserve, &out.Reserve
		*out = new(v1alpha2.PluginSet)
		if err := Convert_config_PluginSet_To_v1alpha2_PluginSet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Reserve = nil
	}
	if in.Permit != nil {
		in, out := &in.Permit, &out.Permit
		*out = new(v1alpha2.PluginSet)
		if err := Convert_config_PluginSet_To_v1alpha2_PluginSet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Permit = nil
	}
	if in.PreBind != nil {
		in, out := &in.PreBind, &out.PreBind
		*out = new(v1alpha2.PluginSet)
		if err := Convert_config_PluginSet_To_v1alpha2_PluginSet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.PreBind = nil
	}
	if in.Bind != nil {
		in, out := &in.Bind, &out.Bind
		*out = new(v1alpha2.PluginSet)
		if err := Convert_config_PluginSet_To_v1alpha2_PluginSet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Bind = nil
	}
	if in.PostBind != nil {
		in, out := &in.PostBind, &out.PostBind
		*out = new(v1alpha2.PluginSet)
		if err := Convert_config_PluginSet_To_v1alpha2_PluginSet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.PostBind = nil
	}
	if in.Unreserve != nil {
		in, out := &in.Unreserve, &out.Unreserve
		*out = new(v1alpha2.PluginSet)
		if err := Convert_config_PluginSet_To_v1alpha2_PluginSet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Unreserve = nil
	}
	return nil
}

// Convert_config_Plugins_To_v1alpha2_Plugins is an autogenerated conversion function.
func Convert_config_Plugins_To_v1alpha2_Plugins(in *config.Plugins, out *v1alpha2.Plugins, s conversion.Scope) error {
	return autoConvert_config_Plugins_To_v1alpha2_Plugins(in, out, s)
}

func autoConvert_v1alpha2_SchedulerAlgorithmSource_To_config_SchedulerAlgorithmSource(in *v1alpha2.SchedulerAlgorithmSource, out *config.SchedulerAlgorithmSource, s conversion.Scope) error {
	out.Policy = (*config.SchedulerPolicySource)(unsafe.Pointer(in.Policy))
	out.Provider = (*string)(unsafe.Pointer(in.Provider))
	return nil
}

// Convert_v1alpha2_SchedulerAlgorithmSource_To_config_SchedulerAlgorithmSource is an autogenerated conversion function.
func Convert_v1alpha2_SchedulerAlgorithmSource_To_config_SchedulerAlgorithmSource(in *v1alpha2.SchedulerAlgorithmSource, out *config.SchedulerAlgorithmSource, s conversion.Scope) error {
	return autoConvert_v1alpha2_SchedulerAlgorithmSource_To_config_SchedulerAlgorithmSource(in, out, s)
}

func autoConvert_config_SchedulerAlgorithmSource_To_v1alpha2_SchedulerAlgorithmSource(in *config.SchedulerAlgorithmSource, out *v1alpha2.SchedulerAlgorithmSource, s conversion.Scope) error {
	out.Policy = (*v1alpha2.SchedulerPolicySource)(unsafe.Pointer(in.Policy))
	out.Provider = (*string)(unsafe.Pointer(in.Provider))
	return nil
}

// Convert_config_SchedulerAlgorithmSource_To_v1alpha2_SchedulerAlgorithmSource is an autogenerated conversion function.
func Convert_config_SchedulerAlgorithmSource_To_v1alpha2_SchedulerAlgorithmSource(in *config.SchedulerAlgorithmSource, out *v1alpha2.SchedulerAlgorithmSource, s conversion.Scope) error {
	return autoConvert_config_SchedulerAlgorithmSource_To_v1alpha2_SchedulerAlgorithmSource(in, out, s)
}

func autoConvert_v1alpha2_SchedulerPolicyConfigMapSource_To_config_SchedulerPolicyConfigMapSource(in *v1alpha2.SchedulerPolicyConfigMapSource, out *config.SchedulerPolicyConfigMapSource, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	return nil
}

// Convert_v1alpha2_SchedulerPolicyConfigMapSource_To_config_SchedulerPolicyConfigMapSource is an autogenerated conversion function.
func Convert_v1alpha2_SchedulerPolicyConfigMapSource_To_config_SchedulerPolicyConfigMapSource(in *v1alpha2.SchedulerPolicyConfigMapSource, out *config.SchedulerPolicyConfigMapSource, s conversion.Scope) error {
	return autoConvert_v1alpha2_SchedulerPolicyConfigMapSource_To_config_SchedulerPolicyConfigMapSource(in, out, s)
}

func autoConvert_config_SchedulerPolicyConfigMapSource_To_v1alpha2_SchedulerPolicyConfigMapSource(in *config.SchedulerPolicyConfigMapSource, out *v1alpha2.SchedulerPolicyConfigMapSource, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	return nil
}

// Convert_config_SchedulerPolicyConfigMapSource_To_v1alpha2_SchedulerPolicyConfigMapSource is an autogenerated conversion function.
func Convert_config_SchedulerPolicyConfigMapSource_To_v1alpha2_SchedulerPolicyConfigMapSource(in *config.SchedulerPolicyConfigMapSource, out *v1alpha2.SchedulerPolicyConfigMapSource, s conversion.Scope) error {
	return autoConvert_config_SchedulerPolicyConfigMapSource_To_v1alpha2_SchedulerPolicyConfigMapSource(in, out, s)
}

func autoConvert_v1alpha2_SchedulerPolicyFileSource_To_config_SchedulerPolicyFileSource(in *v1alpha2.SchedulerPolicyFileSource, out *config.SchedulerPolicyFileSource, s conversion.Scope) error {
	out.Path = in.Path
	return nil
}

// Convert_v1alpha2_SchedulerPolicyFileSource_To_config_SchedulerPolicyFileSource is an autogenerated conversion function.
func Convert_v1alpha2_SchedulerPolicyFileSource_To_config_SchedulerPolicyFileSource(in *v1alpha2.SchedulerPolicyFileSource, out *config.SchedulerPolicyFileSource, s conversion.Scope) error {
	return autoConvert_v1alpha2_SchedulerPolicyFileSource_To_config_SchedulerPolicyFileSource(in, out, s)
}

func autoConvert_config_SchedulerPolicyFileSource_To_v1alpha2_SchedulerPolicyFileSource(in *config.SchedulerPolicyFileSource, out *v1alpha2.SchedulerPolicyFileSource, s conversion.Scope) error {
	out.Path = in.Path
	return nil
}

// Convert_config_SchedulerPolicyFileSource_To_v1alpha2_SchedulerPolicyFileSource is an autogenerated conversion function.
func Convert_config_SchedulerPolicyFileSource_To_v1alpha2_SchedulerPolicyFileSource(in *config.SchedulerPolicyFileSource, out *v1alpha2.SchedulerPolicyFileSource, s conversion.Scope) error {
	return autoConvert_config_SchedulerPolicyFileSource_To_v1alpha2_SchedulerPolicyFileSource(in, out, s)
}

func autoConvert_v1alpha2_SchedulerPolicySource_To_config_SchedulerPolicySource(in *v1alpha2.SchedulerPolicySource, out *config.SchedulerPolicySource, s conversion.Scope) error {
	out.File = (*config.SchedulerPolicyFileSource)(unsafe.Pointer(in.File))
	out.ConfigMap = (*config.SchedulerPolicyConfigMapSource)(unsafe.Pointer(in.ConfigMap))
	return nil
}

// Convert_v1alpha2_SchedulerPolicySource_To_config_SchedulerPolicySource is an autogenerated conversion function.
func Convert_v1alpha2_SchedulerPolicySource_To_config_SchedulerPolicySource(in *v1alpha2.SchedulerPolicySource, out *config.SchedulerPolicySource, s conversion.Scope) error {
	return autoConvert_v1alpha2_SchedulerPolicySource_To_config_SchedulerPolicySource(in, out, s)
}

func autoConvert_config_SchedulerPolicySource_To_v1alpha2_SchedulerPolicySource(in *config.SchedulerPolicySource, out *v1alpha2.SchedulerPolicySource, s conversion.Scope) error {
	out.File = (*v1alpha2.SchedulerPolicyFileSource)(unsafe.Pointer(in.File))
	out.ConfigMap = (*v1alpha2.SchedulerPolicyConfigMapSource)(unsafe.Pointer(in.ConfigMap))
	return nil
}

// Convert_config_SchedulerPolicySource_To_v1alpha2_SchedulerPolicySource is an autogenerated conversion function.
func Convert_config_SchedulerPolicySource_To_v1alpha2_SchedulerPolicySource(in *config.SchedulerPolicySource, out *v1alpha2.SchedulerPolicySource, s conversion.Scope) error {
	return autoConvert_config_SchedulerPolicySource_To_v1alpha2_SchedulerPolicySource(in, out, s)
}
