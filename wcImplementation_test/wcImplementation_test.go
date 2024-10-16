package wcImplementation_test

import (
	"ccwc/wcImplementation"
	"fmt"
	"go.uber.org/mock/gomock"
	"io"
	"reflect"
	"testing"
)

func generateFileReadCalls(mockOsFile *wcImplementation.MockOSFile, numberOfOperations int) []any {
	const evenCycleCallTimes = 3
	const oddCycleCallTimes = 2

	var buff = make([]byte, 4096)
	var longByte = []byte("Hëllo Wôrd!\n")
	var shortByte = []byte("Wôrd\n")
	calls := make([]any, 0)

	for i := 0; i < numberOfOperations; i++ {
		calls = append(calls, mockOsFile.EXPECT().Read(gomock.AssignableToTypeOf(buff)).DoAndReturn(func(p []byte) (int, error) {
			copy(p, longByte)
			return len(longByte), nil
		}).Times(evenCycleCallTimes))
		calls = append(calls, mockOsFile.EXPECT().Read(gomock.AssignableToTypeOf(buff)).DoAndReturn(func(p []byte) (int, error) {
			copy(p, shortByte)
			return len(shortByte), nil
		}).Times(oddCycleCallTimes))
		calls = append(calls, mockOsFile.EXPECT().Read(gomock.AssignableToTypeOf(buff)).Return(0, io.EOF))
		calls = append(calls, mockOsFile.EXPECT().Seek(int64(0), io.SeekStart).Return(int64(0), nil))
	}

	return calls
}

func TestGetFileStateByteSize(t *testing.T) {
	expectedFileStat := []string{"54"}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOsFile := wcImplementation.NewMockOSFile(ctrl)

	calls := generateFileReadCalls(mockOsFile, 1)
	gomock.InOrder(calls...)

	fileStat, _ := wcImplementation.GetFileState(mockOsFile, true, false, false, false)

	if !reflect.DeepEqual(fileStat, expectedFileStat) {
		t.Error(fmt.Sprintf("Checking byte size failed expecting %v but got %v", expectedFileStat, fileStat))
	}
}

func TestGetFileStateNumberOfLines(t *testing.T) {
	expectedFileStat := []string{"5"}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOsFile := wcImplementation.NewMockOSFile(ctrl)

	calls := generateFileReadCalls(mockOsFile, 1)
	gomock.InOrder(calls...)

	fileStat, _ := wcImplementation.GetFileState(mockOsFile, false, true, false, false)

	if !reflect.DeepEqual(fileStat, expectedFileStat) {
		t.Error(fmt.Sprintf("Checking number of lines failed expecting %v but got %v", expectedFileStat, fileStat))
	}
}

func TestGetFileStateNumberOfWords(t *testing.T) {
	expectedFileStat := []string{"8"}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOsFile := wcImplementation.NewMockOSFile(ctrl)

	calls := generateFileReadCalls(mockOsFile, 1)
	gomock.InOrder(calls...)

	fileStat, _ := wcImplementation.GetFileState(mockOsFile, false, false, true, false)

	if !reflect.DeepEqual(fileStat, expectedFileStat) {
		t.Error(fmt.Sprintf("Checking number of words failed expecting %v but got %v", expectedFileStat, fileStat))
	}
}

func TestGetFileStateNumberOfChars(t *testing.T) {
	expectedFileStat := []string{"46"}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOsFile := wcImplementation.NewMockOSFile(ctrl)

	calls := generateFileReadCalls(mockOsFile, 1)
	gomock.InOrder(calls...)

	fileStat, _ := wcImplementation.GetFileState(mockOsFile, false, false, false, true)

	if !reflect.DeepEqual(fileStat, expectedFileStat) {
		t.Error(fmt.Sprintf("Checking number of chars failed expecting %v but got %v", expectedFileStat, fileStat))
	}
}

func TestGetFileStateNumberOfLinesAndWords(t *testing.T) {
	expectedFileStat := []string{"5", "8"}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOsFile := wcImplementation.NewMockOSFile(ctrl)

	calls := generateFileReadCalls(mockOsFile, 2)
	gomock.InOrder(calls...)

	fileStat, _ := wcImplementation.GetFileState(mockOsFile, false, true, true, false)

	if !reflect.DeepEqual(fileStat, expectedFileStat) {
		t.Error(fmt.Sprintf("Checking number of lines and words failed expecting %v but got %v", expectedFileStat, fileStat))
	}
}

func TestGetAllFileState(t *testing.T) {
	expectedFileStat := []string{"54", "5", "8", "46"}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOsFile := wcImplementation.NewMockOSFile(ctrl)

	calls := generateFileReadCalls(mockOsFile, 4)
	gomock.InOrder(calls...)

	fileStat, _ := wcImplementation.GetFileState(mockOsFile, true, true, true, true)

	if !reflect.DeepEqual(fileStat, expectedFileStat) {
		t.Error(fmt.Sprintf("Checking all file stats failed expecting %v but got %v", expectedFileStat, fileStat))
	}
}
