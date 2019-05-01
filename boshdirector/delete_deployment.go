// Copyright (C) 2016-Present Pivotal Software, Inc. All rights reserved.
// This program and the accompanying materials are made available under the terms of the under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

package boshdirector

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

func (c *Client) DeleteDeployment(name, contextID string, force bool, taskReporter *AsyncTaskReporter, logger *log.Logger) (int, error) {
	logger.Printf("deleting deployment %s\n", name)
	d, err := c.Director(taskReporter)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to build director")
	}

	_, found, err := c.GetDeployment(name, logger)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf(`BOSH error when deleting deployment "%s"`, name))
	}
	if !found {
		taskReporter.TaskStarted(0)
		taskReporter.TaskFinished(0, "done")
		return 0, nil
	}

	deployment, err := d.WithContext(contextID).FindDeployment(name)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf(`BOSH error when deleting deployment "%s"`, name))
	}
	go func() {
		err = deployment.Delete(force)
		if err != nil {
			taskReporter.Err <- errors.Wrap(err, fmt.Sprintf("Could not delete deployment %s", name))
		}
	}()

	select {
	case err := <-taskReporter.Err:
		return 0, err
	case id := <-taskReporter.Task:
		return id, nil
	}
}
