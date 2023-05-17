package server

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"homework-5/internal/pkg/repository"
	mock_repository "homework-5/internal/pkg/repository/mocks"
	"net/http"
	"net/url"
	"testing"
)

func Test_getStudent(t *testing.T) {
	ctx := context.Background()
	var id uint64 = 1
	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		//arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := mock_repository.NewMockStudentRepo(ctrl)
		s := server{studentRepo: m}

		req, err := http.NewRequest(http.MethodGet, "?id=1", bytes.NewReader([]byte{}))
		require.NoError(t, err)
		m.EXPECT().GetById(gomock.Any(), id).Return(&repository.Student{ID: 1, Name: "test", Grades: 250, UnivID: 1}, nil)
		//act
		_, status := s.getStudent(ctx, req)
		//assert
		require.Equal(t, http.StatusOK, status)

	})
	t.Run("fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Parallel()
		tt := []struct {
			name    string
			request *url.URL
			isOk    bool
		}{
			{
				"without id",
				&url.URL{RawQuery: "id"},
				false,
			},
			{
				"wrong id",
				&url.URL{RawQuery: "id=asdasd"},
				false,
			},
			{
				"empty",
				&url.URL{RawQuery: ""},
				false,
			},
			{
				"ok",
				&url.URL{RawQuery: "id=0"},
				true,
			},
		}
		for _, tc := range tt {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				id, err := getStudentID(tc.request)
				if !tc.isOk {
					assert.EqualError(t, err, "can't get id")
				} else {
					assert.Equal(t, uint64(0), id)
				}
			})
		}

	})

}

func Test_studentDelete(t *testing.T) {
	ctx := context.Background()
	var id uint64 = 1
	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		//arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := mock_repository.NewMockStudentRepo(ctrl)
		s := server{studentRepo: m}

		req, err := http.NewRequest(http.MethodGet, "?id=1", bytes.NewReader([]byte{}))
		require.NoError(t, err)
		m.EXPECT().Delete(gomock.Any(), id).Return(true, nil)
		//act
		_, status := s.deleteStudent(ctx, req)
		//assert
		require.Equal(t, http.StatusOK, status)

	})
	t.Run("fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Parallel()
		tt := []struct {
			name    string
			request *url.URL
			isOk    bool
		}{
			{
				"without id",
				&url.URL{RawQuery: "id"},
				false,
			},
			{
				"wrong id",
				&url.URL{RawQuery: "id=asdasd"},
				false,
			},
			{
				"empty",
				&url.URL{RawQuery: ""},
				false,
			},
			{
				"ok",
				&url.URL{RawQuery: "id=0"},
				true,
			},
		}
		for _, tc := range tt {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				id, err := getStudentID(tc.request)
				if !tc.isOk {
					assert.EqualError(t, err, "can't get id")
				} else {
					assert.Equal(t, uint64(0), id)
				}
			})
		}

	})

}

func Test_studentUpdate(t *testing.T) {

	ctx := context.Background()
	var data repository.Student
	data.ID = uint64(1)
	data.Name = "meme"
	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		//arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := mock_repository.NewMockStudentRepo(ctrl)

		s := server{studentRepo: m}

		json, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}

		req, err := http.NewRequest(http.MethodPut, "/", bytes.NewBuffer(json))
		require.NoError(t, err)
		m.EXPECT().Update(gomock.Any(), &data).Return(true, nil)
		//act
		_, status := s.updateStudent(ctx, req)
		//assert
		require.Equal(t, http.StatusOK, status)

	})
	t.Run("fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		//t.Parallel()
		tt := []struct {
			name    string
			request string
			isOk    bool
		}{
			{
				"correct",
				"{\"id\":1,\"name\":\"meme\"}",
				true,
			},
			{
				"invalid type",
				"{{\"id\":\"dsds\",\"name\":\"meme\"}",
				false,
			},
			{
				"empty",
				"",
				false,
			},
			{
				"empty fields fields",
				"{\"id\":,\"name\":\"meme\"}",
				false,
			},
		}

		m := mock_repository.NewMockStudentRepo(ctrl)
		s := server{studentRepo: m}
		for _, tc := range tt {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				byteRequest := []byte(tc.request)
				req, _ := http.NewRequest(http.MethodPut, "/", bytes.NewBuffer(byteRequest))
				m.EXPECT().Update(gomock.Any(), &data).Return(true, nil)
				_, status := s.updateStudent(ctx, req)
				if !tc.isOk {

					assert.Equal(t, http.StatusBadRequest, status)
				} else {
					assert.Equal(t, http.StatusOK, status)
				}
			})
		}

	})

}

func Test_studentCreate(t *testing.T) {

	ctx := context.Background()
	var data repository.Student
	data.Name = "meme"
	data.Grades = 250
	data.UnivID = 25
	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		//arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := mock_repository.NewMockStudentRepo(ctrl)

		s := server{studentRepo: m}

		json, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(json))
		require.NoError(t, err)
		m.EXPECT().Add(gomock.Any(), &data).Return(uint64(1), nil)
		//act
		_, status := s.createStudent(ctx, req)
		//assert
		require.Equal(t, http.StatusOK, status)

	})
	t.Run("fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		t.Parallel()
		tt := []struct {
			name    string
			request string
			isOk    bool
		}{
			{
				"correct",
				"{\"name\":\"meme\",\"grades\":250,\"univ_id\":25}",
				true,
			},
			{
				"invalid type",
				"{\"name\":\"test\",\"grades\":\"asd\",\"univ_id\":0}",
				false,
			},
			{
				"empty",
				"",
				false,
			},
			{
				"empty fields fields",
				"{\"nadsdsme\":,\"grasdsdsdes\":\"asd\",\"undsdsdiv_id\":0}",
				false,
			},
		}

		m := mock_repository.NewMockStudentRepo(ctrl)
		s := server{studentRepo: m}
		for _, tc := range tt {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				byteRequest := []byte(tc.request)
				req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(byteRequest))
				m.EXPECT().Add(gomock.Any(), &data).Return(uint64(1), nil)
				_, status := s.createStudent(ctx, req)

				if !tc.isOk {
					//assert.EqualError(t, err, "wrong input type")
					assert.Equal(t, http.StatusBadRequest, status)
				} else {
					//assert.Equal(t, "meme", *studentData.Name)
					assert.Equal(t, http.StatusOK, status)
				}
			})
		}

	})

}

func Test_GetStudentData(t *testing.T) {
	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		tt := []struct {
			name    string
			request string
			isOk    bool
		}{
			{
				"correct",
				"{\"name\":\"meme\",\"grades\":250,\"univ_id\":25}",
				true,
			},
			{
				"invalid type",
				"{\"name\":\"test\",\"grades\":\"asd\",\"univ_id\":0}",
				false,
			},
			{
				"empty",
				"",
				false,
			},
			{
				"empty fields fields",
				"{\"nadsdsme\":,\"grasdsdsdes\":\"asd\",\"undsdsdiv_id\":0}",
				false,
			},
		}
		for _, tc := range tt {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				byteRequest := []byte(tc.request)
				req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(byteRequest))
				id, err := getStudentData(req.Body)
				if !tc.isOk {
					assert.EqualError(t, err, "wrong input type")
				} else {
					assert.Equal(t, "meme", *id.Name)
				}
			})
		}

	})
}

func Test_getUniversity(t *testing.T) {
	ctx := context.Background()
	var id uint64 = 1
	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		//arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := mock_repository.NewMockUniversityRepo(ctrl)
		s := server{universityRepo: m}

		req, err := http.NewRequest(http.MethodGet, "?id=1", bytes.NewReader([]byte{}))
		require.NoError(t, err)
		m.EXPECT().GetById(gomock.Any(), id).Return(&repository.University{ID: 1, Name: "test", Facility: "testFacility"}, nil)
		//act
		_, status := s.getUniversity(ctx, req)
		//assert
		require.Equal(t, http.StatusOK, status)

	})
	t.Run("fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Parallel()
		tt := []struct {
			name    string
			request *url.URL
			isOk    bool
		}{
			{
				"/without id",
				&url.URL{RawQuery: "id"},
				false,
			},
			{
				"wrong id",
				&url.URL{RawQuery: "id=asdasd"},
				false,
			},
			{
				"empty",
				&url.URL{RawQuery: ""},
				false,
			},
			{
				"ok",
				&url.URL{RawQuery: "id=0"},
				true,
			},
		}
		for _, tc := range tt {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				id, err := getUniversityID(tc.request)
				if !tc.isOk {
					assert.EqualError(t, err, "can't get id")
				} else {
					assert.Equal(t, uint64(0), id)
				}
			})
		}

	})

}

func Test_universityDelete(t *testing.T) {
	ctx := context.Background()
	var id uint64 = 1
	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		//arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := mock_repository.NewMockUniversityRepo(ctrl)
		s := server{universityRepo: m}

		req, err := http.NewRequest(http.MethodGet, "?id=1", bytes.NewReader([]byte{}))
		require.NoError(t, err)
		m.EXPECT().Delete(gomock.Any(), id).Return(true, nil)
		//act
		_, status := s.deleteUniversity(ctx, req)
		//assert
		require.Equal(t, http.StatusOK, status)

	})
	t.Run("fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Parallel()
		tt := []struct {
			name    string
			request *url.URL
			isOk    bool
		}{
			{
				"without id",
				&url.URL{RawQuery: "id"},
				false,
			},
			{
				"wrong id",
				&url.URL{RawQuery: "id=asdasd"},
				false,
			},
			{
				"empty",
				&url.URL{RawQuery: ""},
				false,
			},
			{
				"ok",
				&url.URL{RawQuery: "id=0"},
				true,
			},
		}
		for _, tc := range tt {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				id, err := getUniversityID(tc.request)
				if !tc.isOk {
					assert.EqualError(t, err, "can't get id")
				} else {
					assert.Equal(t, uint64(0), id)
				}
			})
		}

	})

}

func Test_universityUpdate(t *testing.T) {

	ctx := context.Background()
	var data repository.University
	data.ID = uint64(1)
	data.Name = "meme"
	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		//arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := mock_repository.NewMockUniversityRepo(ctrl)

		s := server{universityRepo: m}

		json, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}

		req, err := http.NewRequest(http.MethodPut, "/", bytes.NewBuffer(json))
		require.NoError(t, err)
		m.EXPECT().Update(gomock.Any(), &data).Return(true, nil)
		//act
		_, status := s.updateUniversity(ctx, req)
		//assert
		require.Equal(t, http.StatusOK, status)

	})
	t.Run("fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		//t.Parallel()
		tt := []struct {
			name    string
			request string
			isOk    bool
		}{
			{
				"correct",
				"{\"id\":1,\"name\":\"meme\"}",
				true,
			},
			{
				"invalid type",
				"{{\"id\":\"dsds\",\"name\":\"meme\"}",
				false,
			},
			{
				"empty",
				"",
				false,
			},
			{
				"empty fields fields",
				"{\"id\":,\"name\":\"meme\"}",
				false,
			},
		}

		m := mock_repository.NewMockUniversityRepo(ctrl)
		s := server{universityRepo: m}
		for _, tc := range tt {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				byteRequest := []byte(tc.request)
				req, _ := http.NewRequest(http.MethodPut, "/", bytes.NewBuffer(byteRequest))
				m.EXPECT().Update(gomock.Any(), &data).Return(true, nil)
				_, status := s.updateUniversity(ctx, req)
				if !tc.isOk {

					assert.Equal(t, http.StatusBadRequest, status)
				} else {
					assert.Equal(t, http.StatusOK, status)
				}
			})
		}

	})

}

func Test_universityCreate(t *testing.T) {

	ctx := context.Background()
	var data repository.University
	data.Name = "meme"
	data.Facility = "facilityTest"
	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		//arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := mock_repository.NewMockUniversityRepo(ctrl)

		s := server{universityRepo: m}

		json, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(json))
		require.NoError(t, err)
		m.EXPECT().Add(gomock.Any(), &data).Return(int64(1), nil)
		//act
		_, status := s.createUniversity(ctx, req)
		//assert
		require.Equal(t, http.StatusOK, status)

	})
	t.Run("fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		t.Parallel()
		tt := []struct {
			name    string
			request string
			isOk    bool
		}{
			{
				"invalid type",
				"{\"name\":163,\"facility\":\"sdsdsd\"}",
				false,
			},
			{
				"empty",
				"",
				false,
			},
			{
				"empty fields fields",
				"{\"nadsdsme\":,\"grasdsdsdes\":\"asd\"}",
				false,
			},
		}

		m := mock_repository.NewMockUniversityRepo(ctrl)
		s := server{universityRepo: m}
		for _, tc := range tt {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				byteRequest := []byte(tc.request)
				req, _ := http.NewRequest(http.MethodPost, universityPath, bytes.NewBuffer(byteRequest))
				m.EXPECT().Add(gomock.Any(), &data).Return(int64(1), nil)
				_, status := s.createStudent(ctx, req)

				if !tc.isOk {
					assert.Equal(t, http.StatusBadRequest, status)
				} else {
					assert.Equal(t, http.StatusOK, status)
				}
			})
		}

	})

}

func Test_GetUniversityData(t *testing.T) {
	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		tt := []struct {
			name    string
			request string
			isOk    bool
		}{
			{
				"correct",
				"{\"name\":\"meme\",\"facility\":\"test\"}",
				true,
			},
			{
				"invalid type",
				"{\"name\":163,\"facility\":\"test\"}",
				false,
			},
			{
				"empty",
				"",
				false,
			},
			{
				"empty fields fields",
				"{\"name\":\"meme\",\"facility\":}",
				false,
			},
		}
		for _, tc := range tt {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				byteRequest := []byte(tc.request)
				req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(byteRequest))
				id, err := getUniversityData(req.Body)
				if !tc.isOk {
					assert.EqualError(t, err, "wrong input type")
				} else {
					assert.Equal(t, "meme", *id.Name)
				}
			})
		}

	})
}
