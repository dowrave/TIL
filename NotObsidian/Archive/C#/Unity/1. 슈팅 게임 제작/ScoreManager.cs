using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class ScoreManager : MonoBehaviour
{

    public Text currentScoreUI;
    private int CurrentScore;
    public Text bestScoreUI;
    private int bestScore;

    // �̱��� ��ü
    public static ScoreManager Instance = null;

    public int Score
    {
        get
        {
            return CurrentScore;
        }
        set
        {
            // 3. ���� �Ҵ��Ѵ�
            CurrentScore = value;

            // 4. ���� ǥ���ϱ�
            currentScoreUI.text = "��������  " + CurrentScore;

            if (CurrentScore > bestScore)
            {
                bestScore = CurrentScore;
                bestScoreUI.text = "�ְ�����  " + bestScore;

                // ���� �����ϱ�
                PlayerPrefs.SetInt("Best Score", bestScore);
            }
        }
    }

    // �̱��� ��ü�� ���� ������ ������ �ڱ� �ڽ��� �Ҵ���
    void Awake()
    {
        if(Instance == null)
        {
            Instance = this;
        }
    }


    public int GetScore()
    {
        return CurrentScore;
    }
    void Start()
    {
        // 1. �ְ� ���� �ҷ�����
        bestScore = PlayerPrefs.GetInt("Best Score", 0);
        // 2. �ְ� ���� ǥ���ϱ�
        bestScoreUI.text = "�ְ� ���� : " + bestScore;
    }
}
